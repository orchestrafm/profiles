package identity

import (
	"os"

	"github.com/Nerzal/gocloak"
	"github.com/spidernest-go/logger"
)

var idp gocloak.GoCloak
var token *gocloak.JWT

func Handshake() {
	var err error
	idp = gocloak.NewClient(os.Getenv("IDP_ADDR"))
	token, err = idp.LoginAdmin(os.Getenv("IDP_USER"), os.Getenv("IDP_PASS"), "master")
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Handshake with IDP server failed.")
	}
}
