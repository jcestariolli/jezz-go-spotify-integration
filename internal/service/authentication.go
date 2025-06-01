package service

import (
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
)

type AuthService struct {
	authClient client.AuthClient
}

func NewAuthService(authClient client.AuthClient) AuthService {
	return AuthService{
		authClient: authClient,
	}
}

func (s AuthService) AuthenticateApp() (model.AuthSession, error) {
	return s.authClient.Authenticate()
}
