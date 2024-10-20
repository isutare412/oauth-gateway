package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/isutare412/oauth-gateway/internal/core/model"
)

type googleIDTokenClaims struct {
	jwt.RegisteredClaims
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	PicturlURL    string `json:"picture"`
	ProfileURL    string `json:"profile"`
}

func (c *googleIDTokenClaims) ToGoogleIDToken() *model.GoogleIDToken {
	return &model.GoogleIDToken{
		Subject:    c.Subject,
		IssuedAt:   c.IssuedAt.Time,
		ExpiresAt:  c.ExpiresAt.Time,
		Email:      c.Email,
		Name:       c.Name,
		FamilyName: c.FamilyName,
		GivenName:  c.GivenName,
		PictureURL: c.PicturlURL,
	}
}
