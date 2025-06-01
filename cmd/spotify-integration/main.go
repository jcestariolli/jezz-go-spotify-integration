package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/service"
)

func main() {
	var spotifyConfig config.Config
	if loadConfig(&spotifyConfig) != true {
		return
	}
	authService, artistApi := loadServices(spotifyConfig)
	runApp(authService, artistApi)
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

func loadServices(config config.Config) (service.AuthService, client.ArtistsAPIClient) {
	fmt.Println("Loading spotify services...")
	cliConfig := config.Client
	authService := loadAuthService(cliConfig)
	artistApi := loadArtistApi(cliConfig)
	fmt.Printf("✔ Services loaded! :)\n\n")
	return authService, artistApi
}

func loadArtistApi(cliConfig config.CliConfig) client.ArtistsAPIClient {
	return client.NewArtistsAPIClient(cliConfig.BaseUrl)
}

func loadAuthService(cliConfig config.CliConfig) service.AuthService {
	authClient := client.NewCliCredentialsAuth(
		cliConfig.BaseUrl,
		cliConfig.AccountsUrl,
		cliConfig.CliCredentials,
	)
	authService := service.NewAuthService(authClient)
	return authService
}

func runApp(authService service.AuthService, artistApi client.ArtistsAPIClient) {
	authSession, err := authenticateApp(authService)
	if err == true {
		return
	}
	if getArtist(artistApi, authSession) {
		return
	}

	return
}

func getArtist(artistApi client.ArtistsAPIClient, authSession model.AuthSession) bool {
	fmt.Println("Trying to get artist...")
	artist, err := artistApi.GetArtist(authSession.Auth.AccessToken, "7nzSoJISlVJsn7O0yTeMOB?si=1RkwrfE4QWanTYQMdN1pTg")

	if err != nil {
		fmt.Println("✖ Getting artist failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return true
	}
	fmt.Println("✔ Artist obtained! :)")
	if body, err3 := json.Marshal(artist); err3 == nil && body != nil {
		fmt.Printf("╰┈➤%s\n", string(body))
	}
	fmt.Printf("\n")
	return false
}

func authenticateApp(authService service.AuthService) (model.AuthSession, bool) {
	fmt.Println("Trying to authenticate application...")
	authSession, err := authService.AuthenticateApp()
	if err != nil {
		fmt.Println("✖ Authentication failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return authSession, true
	}
	fmt.Println("✔ Authentication succeeded! :)")
	if body, err3 := json.Marshal(authSession); err3 == nil && body != nil {
		fmt.Printf("╰┈➤%s\n", string(body))
	}
	fmt.Printf("\n")
	return authSession, false
}
