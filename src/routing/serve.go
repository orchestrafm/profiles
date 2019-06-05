package routers

import (
	"github.com/spidernest-go/mux"
)

var r *echo.Echo

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`

func ListenAndServe() {
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.GET("/profile/:id", getProfileById)

	r.Start(":5000")
}
