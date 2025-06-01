package client

import "jezz-go-spotify-integration/internal/model"

type AuthClient interface {
	Authenticate() (model.AuthSession, error)
}
