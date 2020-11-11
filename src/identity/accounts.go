package identity

import (
	"os"

	"github.com/Nerzal/gocloak"
)

func NewAccount(username, email, password string) (string, error) {
	user := gocloak.User{
		Email:     email,
		Enabled:   true,
		Username:  username,
		FirstName: username,
	}

	uuid, err := idp.CreateUser(token.AccessToken, os.Getenv("IDP_REALM"), user)
	if err != nil {
		return "", err
	}

	//HACK: hopefully future gocloak versions let me set
	//      the password directly on the type struct
	err = idp.SetPassword(token.AccessToken, uuid, os.Getenv("IDP_REALM"), password, false)

	return uuid, err
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

func GetGroups(uuid string) ([]*gocloak.UserGroup, error) {
	return idp.GetUserGroups(token.AccessToken, os.Getenv("IDP_REALM"), uuid)
}

func RefreshToken(ref string) (*gocloak.JWT, error) {
	return idp.RefreshToken(
		ref,
		os.Getenv("OIDC_CLIENT_ID"),
		os.Getenv("OIDC_CLIENT_SECRET"),
		os.Getenv("IDP_REALM"))
}
