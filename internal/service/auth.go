package service

import (
	"errors"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/model"
)

type SpotifyAuthService struct {
	appAuth  *model.Authentication
	authFlow auth.AuthenticationFlow
}

func NewSpotifyAuthService(
	authFlow auth.AuthenticationFlow,
) (*SpotifyAuthService, error) {
	authentication, err := authFlow.Authenticate()
	if err != nil {
		return nil, err
	}
	return &SpotifyAuthService{
		appAuth:  authentication,
		authFlow: authFlow,
	}, nil
}

func (s *SpotifyAuthService) ExecuteWithAuthentication(fn FnWithAuthentication) (any, error) {
	t, err := s.authAndExecute(false, fn)
	if err != nil {
		apiErr := commons.ResourceError{}
		if errors.As(err, &apiErr) && apiErr.Status == 401 || apiErr.Status == 403 {
			return s.authAndExecute(true, fn)
		}
	}
	return t, err
}

func (s *SpotifyAuthService) authenticate() error {
	authSession, err := s.authFlow.Authenticate()
	if err != nil {
		return err
	}
	s.appAuth = authSession
	return nil
}

func (s *SpotifyAuthService) authAndExecute(forceAuth bool, fn FnWithAuthentication) (any, error) {
	if forceAuth || s.appAuth == nil {
		err := s.authenticate()
		if err != nil {
			return nil, err
		}
	}
	return fn(s.appAuth.AccessToken)
}
