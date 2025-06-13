package main

import (
	_ "embed"
	"fmt"
	"jezz-go-spotify-integration/cmd/spotify-cli/sample"
	"jezz-go-spotify-integration/internal"
	"jezz-go-spotify-integration/internal/albums"
	"jezz-go-spotify-integration/internal/artists"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/tracks"
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
	authService := loadAuthService(appCfg, cliCredCfg)
	artistsSvc, albumSvc, tracksSvc := loadServices(appCfg, authService)

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

func loadAuthService(appCfg config.AppConfig, cliCredCfg config.CliCredentials) *auth.Service {
	fmt.Println("Loading auth service...")
	authService, err := auth.NewService(
		appCfg.Client.AccountsUrl,
		cliCredCfg,
	)
	if err != nil {
		fmt.Println("✖ Auth service loading failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return nil
	}
	fmt.Printf("✔ Auth service loaded! :)\n\n")
	return authService
}

func loadServices(cfg config.AppConfig, authService *auth.Service) (internal.ArtistsService, internal.AlbumsService, internal.TracksService) {
	cliConfig := cfg.Client
	artistsSvc := loadArtistsService(cliConfig, authService)
	albumsSvc := loadAlbumsService(cliConfig, authService)
	tracksSvc := loadTracksService(cliConfig, authService)
	return artistsSvc, albumsSvc, tracksSvc
}

func loadArtistsService(cliConfig config.CliConfig, authService *auth.Service) internal.ArtistsService {
	fmt.Println("Loading artists service...")
	artistsSvc := artists.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Artist service loaded! :)\n\n")
	return artistsSvc
}

func loadAlbumsService(cliConfig config.CliConfig, authService *auth.Service) internal.AlbumsService {
	fmt.Println("Loading albums service...")
	albumsSvc := albums.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Album service loaded! :)\n\n")
	return albumsSvc
}

func loadTracksService(cliConfig config.CliConfig, authService *auth.Service) internal.TracksService {
	fmt.Println("Loading tracks service...")
	tracksSvc := tracks.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Track service loaded! :)\n\n")
	return tracksSvc
}
