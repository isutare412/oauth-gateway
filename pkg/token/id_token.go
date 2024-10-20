package token

import "github.com/golang-jwt/jwt/v5"

type IDToken struct {
	jwt.RegisteredClaims

	UserID string
}
