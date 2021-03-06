package routers

import (
	"net/http"

	"github.com/orchestrafm/profiles/src/database"
	"github.com/orchestrafm/profiles/src/identity"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createProfile(c echo.Context) error {
	// Validate Data
	reg := new(database.Registration)
	if err := c.Bind(reg); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed registration form.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Registration form data was invalid or malformed."})
	}

	// Burn Invite Code and reject if already burned
	err := database.BurnInvite(reg.InviteCode)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Invite Code is invalid or already used."})
	}

	// Setup Profile
	uuid, err := identity.NewAccount(reg.Username, reg.Email, reg.Password)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Identity Provider refused to create a new user.")

		err = database.UnburnInvite(reg.InviteCode)
		if err != nil {
			logger.Error().
				Err(err).
				Msg("Invite Code could not be unburned.")
		}

		return c.JSON(http.StatusInternalServerError, &struct {
			Message string
		}{
			Message: "Identity Server failed to create the account."})
	}
	p := new(database.Profile)
	p.UUID = uuid
	err = p.New()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Profile was not inserted into database.")

		err = database.UnburnInvite(reg.InviteCode)
		if err != nil {
			logger.Error().
				Err(err).
				Msg("Invite Code could not be unburned.")
		}

		err = identity.DeleteAccount(uuid)
		if err != nil {
			logger.Error().
				Err(err).
				Msg("User also wasn't removed from IDP.")
		}

		return c.JSON(http.StatusInternalServerError, &struct {
			Message string
		}{
			Message: "Database could not be reached."})
	}

	return c.JSON(http.StatusOK, p)
}

func loginProfile(c echo.Context) error {
	//TODO: this function should use HTTP Basic Auth instead

	// form binding
	lgn := new(database.Login)
	if err := c.Bind(lgn); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed login form.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Login form data was invalid or malformed."})
	}

	// login profile
	jwt, err := identity.LoginAccount(lgn.Username, lgn.Password)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("User failed to login.")

		return c.JSON(http.StatusForbidden, &struct {
			Message string
		}{
			Message: "Incorrect username or password."})
	}

	return c.JSON(http.StatusAccepted, &struct {
		RefreshToken string `json:"refresh"`
		BearerToken  string `json:"bearer"`
	}{
		RefreshToken: jwt.RefreshToken,
		BearerToken:  jwt.AccessToken,
	})
}

func refreshAuth(c echo.Context) error {
	type ref struct {
		Value string `json:"refresh_token"`
	}

	// Bind Data
	tkn := new(ref)
	if err := c.Bind(tkn); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed token.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Refresh token was invalid or malformed."})
	}

	// Refresh Authorization
	if jwt, err := identity.RefreshToken(tkn.Value); err != nil {
		logger.Error().
			Err(err).
			Msg("User attempted to refresh their authorization.")

		return c.JSON(http.StatusInternalServerError, &struct {
			Message string
		}{
			Message: "Identity server rejected the refresh token."})
	} else {
		return c.JSON(http.StatusAccepted, &struct {
			RefreshToken string `json:"refresh"`
			BearerToken  string `json:"bearer"`
		}{
			RefreshToken: jwt.RefreshToken,
			BearerToken:  jwt.AccessToken,
		})
	}
}
