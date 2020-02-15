package routers

import (
	"net/http"

	"github.com/orchestrafm/profiles/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func joinMailingList(c echo.Context) error {
	rq := new(database.ReqList)
	if err := c.Bind(rq); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed email.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Email was invalid or malformed."})
	}

	err := rq.New()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Email did not get submitted to the database."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "OK."})
}
