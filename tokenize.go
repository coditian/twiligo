package twiligo

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtCommClaims struct {
	jwt.StandardClaims
	Scope string
}

func CapabilityToken(appSid, accountSid, authToken, client string) (string, error) {
	claims := &jwtCommClaims{
		Scope: "scope:client:incoming?clientName=" + client + " scope:client:outgoing?appSid=" + appSid + "&clientName=" + client,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    accountSid,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(authToken))
}
