package identity

import (
	"os"

	"github.com/Nerzal/gocloak"
)

func NewAccount(username, email, password string) (string, error) {
	user := gocloak.User{
		Email:    email,
		Enabled:  true,
		Username: username,
	}

	return idp.CreateUser(token.AccessToken, os.Getenv("IDP_REALM"), user)
}

func DeleteAccount(uuid string) error {
	return idp.DeleteUser(token.AccessToken, os.Getenv("IDP_REALM"), uuid)
}

func LoginAccount(username, password string) (*gocloak.JWT, error) {
	return idp.Login(
		os.Getenv("OIDC_CLIENT_ID"),
		os.Getenv("OIDC_CLIENT_SECRET"),
		os.Getenv("IDP_REALM"),
		username,
		password)
}

func GetAccount(uuid string) (*gocloak.User, error) {
	return idp.GetUserByID(token.AccessToken,
		os.Getenv("IDP_REALM"),
		uuid)
}
