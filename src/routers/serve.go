package routers

import (
	"github.com/labstack/echo"
)

var r *echo.Echo

func ListenAndServe() {
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.GET("/profile/:id")

	r.Start(":5000")
}
