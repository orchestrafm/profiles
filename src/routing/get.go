package routers

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/profiles/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
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

	return c.JSON(http.StatusOK, &pf)
}
