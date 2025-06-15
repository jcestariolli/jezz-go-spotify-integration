package main

import (
	_ "embed"
	"fmt"
	"jezz-go-spotify-integration/cmd/spotify-cli/sample"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/service"
)

//go:embed config/config.yml
var appConfigData []byte

//go:embed config/spotify_client_credentials.yml
var spotifyCliCredentialsData []byte

func main() {
	appCfg, cliCredCfg, err := loadConfigs()
	if err != nil {
		return
	}
	httpAPIClient := loadHTTPClients()
	authService := loadAuthService(appCfg, cliCredCfg)
	artistsSvc, albumSvc, tracksSvc := loadServices(appCfg, httpAPIClient, authService)

	sample.RunAppSampleCalls(artistsSvc, albumSvc, tracksSvc)
}

func loadConfigs() (config.AppConfig, config.CliCredentials, error) {
	fmt.Println("Loading app configs...")
	appCfgLoader := NewAppConfigLoader()
	appCfg, err := appCfgLoader.Load(appConfigData)
	if err != nil {
		fmt.Println("✖ Error loading configs :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return config.AppConfig{}, config.CliCredentials{}, err
	}
	fmt.Printf("✔ App configs loaded! :)\n\n")

	fmt.Println("Loading client credentials configs...")
	cliCredCfgLoader := NewCliCredentialsLoader()
	cliCredCfg, err := cliCredCfgLoader.Load(spotifyCliCredentialsData)
	if err != nil {
		fmt.Println("✖ Error loading client credentials configs :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return config.AppConfig{}, config.CliCredentials{}, err
	}
	fmt.Printf("✔ Client credentials configs loaded! :)\n\n")

	return appCfg, cliCredCfg, nil
}

func NewAppConfigLoader() config.Loader[config.AppConfig] {
	return config.AppConfigLoader{}
}

func NewCliCredentialsLoader() config.Loader[config.CliCredentials] {
	return config.CliCredentialsConfigLoader{}
}

func NewHTTPApiClient() client.HTTPApiClient {
	return client.NewCustomHTTPApiClient()
}

func loadHTTPClients() client.HTTPApiClient {
	fmt.Println("Loading HTTP API client...")
	httpClient := NewHTTPApiClient()
	fmt.Printf("✔ HTTP API client loaded! :)\n\n")
	return httpClient
}

func loadAuthService(appCfg config.AppConfig, cliCredCfg config.CliCredentials) *service.SpotifyAuthService {
	fmt.Println("Loading auth service...")
	credentialsFlow := auth.NewCliCredentialsFlow(appCfg.Client.AccountsURL, cliCredCfg.ID, cliCredCfg.Secret)
	authService, err := service.NewSpotifyAuthService(credentialsFlow)
	if err != nil {
		fmt.Println("✖ Auth service loading failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return nil
	}
	fmt.Printf("✔ Auth service loaded! :)\n\n")
	return authService
}

func loadServices(cfg config.AppConfig, httpAPIClient client.HTTPApiClient, authService *service.SpotifyAuthService) (service.ArtistsService, service.AlbumsService, service.TracksService) {
	cliConfig := cfg.Client
	artistsSvc := loadArtistsService(cliConfig, httpAPIClient, authService)
	albumsSvc := loadAlbumsService(cliConfig, httpAPIClient, authService)
	tracksSvc := loadTracksService(cliConfig, httpAPIClient, authService)
	return artistsSvc, albumsSvc, tracksSvc
}

func loadArtistsService(cliConfig config.CliConfig, httpAPIClient client.HTTPApiClient, authService *service.SpotifyAuthService) service.ArtistsService {
	fmt.Println("Loading artists service...")
	artistsSvc := service.NewSpotifyArtistsService(
		cliConfig.BaseURL,
		httpAPIClient,
		authService,
	)
	fmt.Printf("✔ Artist service loaded! :)\n\n")
	return artistsSvc
}

func loadAlbumsService(cliConfig config.CliConfig, httpAPIClient client.HTTPApiClient, authService *service.SpotifyAuthService) service.AlbumsService {
	fmt.Println("Loading albums service...")
	albumsSvc := service.NewSpotifyAlbumsService(
		cliConfig.BaseURL,
		httpAPIClient,
		authService,
	)
	fmt.Printf("✔ Album service loaded! :)\n\n")
	return albumsSvc
}

func loadTracksService(cliConfig config.CliConfig, httpAPIClient client.HTTPApiClient, authService *service.SpotifyAuthService) service.TracksService {
	fmt.Println("Loading tracks service...")
	tracksSvc := service.NewSpotifyTracksService(
		cliConfig.BaseURL,
		httpAPIClient,
		authService,
	)
	fmt.Printf("✔ Track service loaded! :)\n\n")
	return tracksSvc
}
