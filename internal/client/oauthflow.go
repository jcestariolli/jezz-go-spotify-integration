package client

import "jezz-go-spotify-integration/internal/model"

type OAuthFlow interface {
	Authenticate() (*model.OAuthResponse, error)
}
