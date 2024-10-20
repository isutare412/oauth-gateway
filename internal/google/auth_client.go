package google

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/isutare412/oauth-gateaway/internal/core/model"
)

type AuthClient struct {
	httpClient *http.Client

	tokenEndpoint     string
	oAuthClientID     string
	oAuthClientSecret string
}

func NewAuthClient(cfg AuthClientConfig) *AuthClient {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100

	return &AuthClient{
		httpClient:        &http.Client{Transport: transport},
		tokenEndpoint:     cfg.TokenEndpoint,
		oAuthClientID:     cfg.OAuthClientID,
		oAuthClientSecret: cfg.OAuthClientSecret,
	}
}

func (c *AuthClient) ExchangeAuthorizationCode(
	ctx context.Context,
	code, redirectURI string,
) (model.GoogleTokenResponse, error) {
	body := url.Values{}
	body.Add("client_id", c.oAuthClientID)
	body.Add("client_secret", c.oAuthClientSecret)
	body.Add("grant_type", "authorization_code")
	body.Add("code", code)
	body.Add("redirect_uri", redirectURI)
	encodedBody := body.Encode()

	req, err := http.NewRequestWithContext(ctx, "POST", c.tokenEndpoint, strings.NewReader(encodedBody))
	if err != nil {
		return model.GoogleTokenResponse{}, fmt.Errorf("building request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rawResp, err := c.httpClient.Do(req)
	switch {
	case err != nil:
		return model.GoogleTokenResponse{}, fmt.Errorf("doing HTTP request: %w", err)
	case rawResp.StatusCode >= 400:
		return model.GoogleTokenResponse{}, fmt.Errorf("got error from Google token API; statusCode(%d)", rawResp.StatusCode)
	}
	defer rawResp.Body.Close()

	var resp googleOAuthTokens
	if err := json.NewDecoder(rawResp.Body).Decode(&resp); err != nil {
		return model.GoogleTokenResponse{}, fmt.Errorf("decoding auth code response: %w", err)
	}

	return resp.ToTokenResponse(), nil
}
