package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	oidc "github.com/coreos/go-oidc"
	"github.com/orchestrafm/profiles/src/database"
	"github.com/orchestrafm/profiles/src/identity"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
	"golang.org/x/oauth2"
)

func getProfileById(c echo.Context) error {
	// Check for valid real numbers
	i, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed id parameter (%s) was not a valid number", c.Param("id"))

		return c.JSON(http.StatusBadRequest, nil)
	}

	// Get profile information
	err, pf := database.SelectProfileById(i)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Profile specified either does not exist or requesting user is unauthorized.")

		c.JSON(http.StatusNotFound, ErrGeneric)
	}

	// Get OIDC Account Info
	acc, err := identity.GetAccount(pf.UUID)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Profile did not have a proper UUID (%s).", pf.UUID)

		return c.JSON(http.StatusBadRequest, nil)
	}
	pf.Username = acc.Username
	pf.UUID = ""

	return c.JSON(http.StatusOK, &pf)
}

func getOIDCLogin(c echo.Context) error {
	StateLock.Lock()
	defer StateLock.Unlock()
	state := identity.GetRandomString(16)
	OAuthStates.Add(state)
	return c.Redirect(http.StatusFound, identity.OAuth2.AuthCodeURL(state, oidc.Nonce(identity.Nonce)))
}

func getOIDCRedirect(c echo.Context) error {
	// match states
	ii := -1
	err := func() error {
		StateLock.RLock()
		defer StateLock.RUnlock()
		if i, val := OAuthStates.Find(func(index int, value interface{}) bool {
			ii = index
			switch value.(string) {
			case c.QueryParams().Get("state"):
				return true
			default:
				return false
			}
		}); i == -1 || val == nil {
			// state did not match
			return c.JSON(http.StatusNotFound, ErrGeneric)
		}
		return nil
	}()
	if err != nil {
		return err
	}

	// remove the state from the list
	StateLock.Lock()
	OAuthStates.Remove(ii)
	StateLock.Unlock()

	tkn, err := identity.OAuth2.Exchange(identity.Context, c.QueryParams().Get("code"))
	if err != nil {
		// failed to exchange token
		return c.JSON(http.StatusNotFound, ErrGeneric)
	}

	raw, ok := tkn.Extra("id_token").(string)
	if !ok {
		// no id_token field in bearer token
		return c.JSON(http.StatusNotFound, ErrGeneric)
	}

	id, err := identity.NonceEnabledVerifier.Verify(identity.Context, raw)
	if err != nil {
		// nonce verification failed
		return c.JSON(http.StatusNotFound, ErrGeneric)
	}
	if id.Nonce != identity.Nonce {
		// nonces do not match and is invalid
		return c.JSON(http.StatusNotFound, ErrGeneric)
	}

	resp := struct {
		Token  *oauth2.Token
		Claims *json.RawMessage
	}{tkn, new(json.RawMessage)}

	if err := id.Claims(&resp.Claims); err != nil {
		return c.JSON(http.StatusNotFound, ErrGeneric)
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrGeneric)
	}

	return c.JSONBlob(http.StatusOK, data)
}
