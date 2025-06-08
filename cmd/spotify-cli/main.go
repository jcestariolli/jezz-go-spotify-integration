package main

import (
	_ "embed"
	"fmt"
	"jezz-go-spotify-integration/cmd/spotify-cli/sample"
	"jezz-go-spotify-integration/internal/albums"
	"jezz-go-spotify-integration/internal/artists"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/tracks"
)

func main() {
	var spotifyConfig config.Config
	if !loadConfig(&spotifyConfig) {
		return
	}
	authService := loadAuthService(spotifyConfig)
	artistsSvc, albumSvc, tracksSvc := loadServices(spotifyConfig, authService)

	sample.RunAppSampleCalls(*artistsSvc, *albumSvc, *tracksSvc)
}

func loadConfig(configPtr *config.Config) bool {
	fmt.Println("Loading spotifyConfig configs...")
	spotifyConfig, err := config.Load()
	if err != nil {
		fmt.Println("✖ Error loading configs :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return false
	}
	*configPtr = spotifyConfig
	fmt.Printf("✔ Configs loaded! :)\n\n")
	return true
}

func loadAuthService(cfg config.Config) *auth.Service {
	fmt.Println("Loading auth service...")
	authService, err := auth.NewService(
		cfg.Client.AccountsUrl,
		cfg.Client.CliCredentials,
	)
	if err != nil {
		fmt.Println("✖ Auth service loading failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return nil
	}
	fmt.Printf("✔ Auth service loaded! :)\n\n")
	return authService
}

func loadServices(cfg config.Config, authService *auth.Service) (*artists.Service, *albums.Service, *tracks.Service) {
	cliConfig := cfg.Client
	artistsSvc := loadArtistsService(cliConfig, authService)
	albumsSvc := loadAlbumsService(cliConfig, authService)
	tracksSvc := loadTracksService(cliConfig, authService)
	return artistsSvc, albumsSvc, tracksSvc
}

func loadArtistsService(cliConfig config.CliConfig, authService *auth.Service) *artists.Service {
	fmt.Println("Loading artists service...")
	artistsSvc := artists.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Artist service loaded! :)\n\n")
	return artistsSvc
}

func loadAlbumsService(cliConfig config.CliConfig, authService *auth.Service) *albums.Service {
	fmt.Println("Loading albums service...")
	albumsSvc := albums.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Album service loaded! :)\n\n")
	return albumsSvc
}

func loadTracksService(cliConfig config.CliConfig, authService *auth.Service) *tracks.Service {
	fmt.Println("Loading tracks service...")
	tracksSvc := tracks.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Track service loaded! :)\n\n")
	return tracksSvc
}
