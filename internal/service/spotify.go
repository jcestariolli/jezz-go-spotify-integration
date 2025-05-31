package service

import (
	"fmt"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/client"
)

type SpotifyService struct {
	authClient client.SpotifyClient
}

func NewSpotifyService(
	authClient client.SpotifyClient,
) SpotifyService {
	return SpotifyService{
		authClient: authClient,
	}
}

func (s *SpotifyService) AuthenticateWithClientCredentials() (auth.AuthorizationHeader, error) {
	accessToken, err := s.getAccessToken()
	if err != nil {
		fmt.Println("Error generating Access Token Header")
		return auth.AuthorizationHeader{}, err
	}
	return auth.AuthorizationHeader{
		Key:   "Authorization",
		Value: auth.BearerAuthorizationType.String() + " " + accessToken.String(),
	}, nil
}

func (s *SpotifyService) getAccessToken() (auth.AccessToken, error) {
	return s.authClient.AuthenticateWithClientCredentials()
}
