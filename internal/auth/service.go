package auth

import (
	"errors"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/model"
)

type Service struct {
	cliCredentials config.CliCredentials
	appAuth        *model.Authentication
	authFlow       CliCredentialsFlow
}

func NewService(
	accountUrl string,
	cliCredentials config.CliCredentials,
) (*Service, error) {
	credentialsFlow := NewCliCredentialsFlow(accountUrl, cliCredentials.Id, cliCredentials.Secret)
	authentication, err := credentialsFlow.Authenticate()
	if err != nil {
		return nil, err
	}
	return &Service{
		cliCredentials: cliCredentials,
		appAuth:        authentication,
		authFlow:       credentialsFlow,
	}, nil
}

func (c *Service) authenticateApp() error {
	authSession, err := c.authFlow.Authenticate()
	if err != nil {
		return err
	}
	c.appAuth = authSession
	return nil
}

func (c *Service) getAppAuth() *model.Authentication {
	return c.appAuth
}

func (c *Service) GetAppAccessToken() model.AccessToken {
	return c.appAuth.AccessToken
}

func ExecuteWithAuthRetry[T any](c *Service, fn func() (T, error)) (T, error) {
	t, err := authAndExecute(c, false, fn)
	if err != nil {
		apiErr := commons.ResourceError{}
		if errors.As(err, &apiErr) && apiErr.Status == 401 || apiErr.Status == 403 {
			return authAndExecute(c, true, fn)
		}
	}
	return t, err
}

func authAndExecute[T any](c *Service, forceAuth bool, fn func() (T, error)) (T, error) {
	var zero T
	if forceAuth || c.appAuth == nil {
		err := c.authenticateApp()
		if err != nil {
			return zero, err
		}
	}
	return fn()
}
