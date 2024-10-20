package port

import (
	"context"

	"github.com/isutare412/oauth-gateaway/internal/core/model"
)

type GoogleAuthClient interface {
	ExchangeAuthorizationCode(ctx context.Context, code, redirectURI string) (model.GoogleTokenResponse, error)
}
