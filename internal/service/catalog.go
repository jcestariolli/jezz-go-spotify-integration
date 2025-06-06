package service

import (
	"errors"
	"jezz-go-spotify-integration/internal/artists"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/config"
)

type CatalogService struct {
	cliCredentials  config.CliCredentials
	appAuth         *auth.Authentication
	authFlow        auth.CliCredentialsFlow
	artistsResource artists.Resource
}

func NewCatalogService(
	baseUrl string,
	accountUrl string,
	cliCredentials config.CliCredentials,
) (*CatalogService, error) {
	authCli := auth.NewCliCredentialsFlow(accountUrl, cliCredentials.Id, cliCredentials.Secret)
	authSession, err := authCli.Authenticate()
	if err != nil {
		return nil, err
	}
	return &CatalogService{
		cliCredentials:  cliCredentials,
		appAuth:         authSession,
		authFlow:        authCli,
		artistsResource: artists.NewResource(baseUrl),
	}, nil
}

func (c *CatalogService) authenticateApp() error {
	authSession, err := c.authFlow.Authenticate()
	if err != nil {
		return err
	}
	c.appAuth = authSession
	return nil
}

func (c *CatalogService) GetArtist(artistId string) (artists.Artist, error) {
	getArtistFn := func() (artists.Artist, error) {
		return c.artistsResource.Get(c.appAuth.AccessToken, artistId)
	}
	return executeWithAuthRetry(c, getArtistFn)
}

func (c *CatalogService) GetArtists(artistIds ...string) (artists.Artists, error) {
	getArtistsFn := func() (artists.Artists, error) {
		return c.artistsResource.GetBatch(c.appAuth.AccessToken, artistIds...)
	}
	return executeWithAuthRetry(c, getArtistsFn)
}

func executeWithAuthRetry[T any](c *CatalogService, fn func() (T, error)) (T, error) {
	t, err := authAndExecute(c, false, fn)
	if err != nil {
		apiErr := commons.ResourceError{}
		if errors.As(err, &apiErr) && apiErr.Status == 401 || apiErr.Status == 403 {
			return authAndExecute(c, true, fn)
		}
	}
	return t, err
}

func authAndExecute[T any](c *CatalogService, forceAuth bool, fn func() (T, error)) (T, error) {
	var zero T
	if forceAuth || c.appAuth == nil {
		err := c.authenticateApp()
		if err != nil {
			return zero, err
		}
	}
	return fn()
}
