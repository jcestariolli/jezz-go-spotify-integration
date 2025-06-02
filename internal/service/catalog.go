package service

import (
	"errors"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
)

type CatalogService struct {
	cliCredentials model.CliCredentials
	appAuthSession *model.AuthSession
	authCli        client.CliCredentialsAuthClient
	artistsCli     client.ArtistsResourceCli
}

func NewCatalogService(
	baseUrl string,
	accountUrl string,
	cliCredentials model.CliCredentials,
) (*CatalogService, error) {
	authCli := client.NewCliCredentialsAuthClient(baseUrl, accountUrl, cliCredentials)
	authSession, err := authCli.Authenticate()
	if err != nil {
		return nil, err
	}
	return &CatalogService{
		cliCredentials: cliCredentials,
		appAuthSession: authSession,
		authCli:        authCli,
		artistsCli:     client.NewArtistsResourceCli(baseUrl),
	}, nil
}

func (c *CatalogService) authenticateApp() error {
	authSession, err := c.authCli.Authenticate()
	if err != nil {
		return err
	}
	c.appAuthSession = authSession
	return nil
}

func (c *CatalogService) GetArtist(artistId string) (model.Artist, error) {
	getArtistFn := func() (model.Artist, error) {
		return c.artistsCli.GetArtist(c.appAuthSession.Auth.AccessToken, artistId)
	}

	artist, err := authenticateAppAndExecute(c, getArtistFn)
	if err != nil {
		apiErr := model.ApiError{}
		if errors.As(err, &apiErr) == true && apiErr.Status == 401 || apiErr.Status == 403 {
			return authenticateAppAndExecute(c, getArtistFn)
		}
		return artist, err
	}
	return artist, nil

}

func authenticateAppAndExecute[T any](c *CatalogService, fn func() (T, error)) (T, error) {
	var zero T
	if c.appAuthSession == nil {
		err := c.authenticateApp()
		if err != nil {
			return zero, err
		}
	}
	return fn()
}
