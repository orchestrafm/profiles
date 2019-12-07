package routers

import (
	"sync"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/orchestrafm/profiles/src/identity"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

var (
	r *echo.Echo

	OAuthStates *arraylist.List
	StateLock   *sync.RWMutex
)

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`

func ListenAndServe() {
	// Initalize stuff for OAuth2 and OpenID Connect
	err := identity.InitRandomPool()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Randomness pool could not be filled, entropy on the current system might be low.")
	}

	StateLock = &sync.RWMutex{}
	OAuthStates = arraylist.New()

	// Start serving API routes
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.GET("/oidc/authorize", getOIDCLogin)
	v0.GET("/oidc/callback", getOIDCRedirect)
	v0.GET("/profile/:id", getProfileById)
	v0.POST("/profile", createProfile)

	r.Start(":5000")
}
