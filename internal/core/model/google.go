package model

import "time"

type GoogleTokenResponse struct {
	AccessToken    string
	AccessTokenTTL time.Duration
	IDToken        string
	Scope          string
	TokenType      string
}

// GoogleIDToken represents fields of google ID token.
// Ref: https://developers.google.com/identity/openid-connect/openid-connect#an-id-tokens-payload.
type GoogleIDToken struct {
	// ID of google account.
	Subject   string
	IssuedAt  time.Time
	ExpiresAt time.Time
	Email     string
	// The user's full name, in a displayable form.
	Name       string
	FamilyName string
	GivenName  string
	PictureURL string
}
