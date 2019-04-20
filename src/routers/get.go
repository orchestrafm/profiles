package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

func getProfile(c echo.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
