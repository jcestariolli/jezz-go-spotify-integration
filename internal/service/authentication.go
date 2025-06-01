package service

import (
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
)

type AuthService struct {
	oAuthClient client.OAuthFlow
}

func NewAuthService(oAuthClient client.OAuthFlow) AuthService {
	return AuthService{
		oAuthClient: oAuthClient,
	}
}

func (s AuthService) Authenticate() (*model.OAuthResponse, error) {
	return s.oAuthClient.Authenticate()
}
