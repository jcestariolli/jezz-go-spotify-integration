package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/artist"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/config"
)

func main() {
	var spotifyConfig config.Config
	if !loadConfig(&spotifyConfig) {
		return
	}
	authService := loadAuthService(spotifyConfig)
	artistService := loadServices(spotifyConfig, authService)

	runAppFixedCalls(*artistService)
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

func loadServices(cfg config.Config, authService *auth.Service) *artist.Service {
	cliConfig := cfg.Client
	catalogService := loadArtistService(cliConfig, authService)
	return catalogService
}

func loadArtistService(cliConfig config.CliConfig, authService *auth.Service) *artist.Service {
	fmt.Println("Loading artist service...")
	artistService := artist.NewService(
		cliConfig.BaseUrl,
		authService,
	)
	fmt.Printf("✔ Artist service loaded! :)\n\n")
	return artistService
}

func runAppFixedCalls(catalogService artist.Service) {
	getArtist(catalogService, "7nzSoJISlVJsn7O0yTeMOB")
	getMultipleArtists(catalogService, "4DFhHyjvGYa9wxdHUjtDkc", "4lgrzShsg2FLA89UM2fdO5")
}

func getArtist(svc artist.Service, artistId string) {
	fmt.Println("Trying to get an artist...")

	artistResponse, err := svc.GetArtist(artistId)
	if err != nil {
		fmt.Println("✖ Getting artist failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(artistResponse); err3 == nil && body != nil {
		fmt.Println("✔ Artist obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting artist failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting artist failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getMultipleArtists(catalogService artist.Service, artistIds ...string) {
	fmt.Println("Trying to get multiple artist...")

	artists, err := catalogService.GetArtists(artistIds...)
	if err != nil {
		fmt.Println("✖ Getting multiple artist failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(artists); err3 == nil && body != nil {
		fmt.Println("✔ Artists obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple artist failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple artist failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}
