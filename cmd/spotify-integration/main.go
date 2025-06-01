package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/app"
	"jezz-go-spotify-integration/internal/client"
)

func main() {
	var config app.Config
	if loadConfig(&config) != true {
		fmt.Println(" ----> Ending program.")
		return
	}
	spotifyClient := loadSpotifyClient(config)

	if runApp(spotifyClient) != true {
		fmt.Println(" ----> Ending program.")
		return
	}
}

func loadConfig(configPtr *app.Config) bool {
	fmt.Println("Loading app configs...")
	config, err := app.Load()
	if err != nil {
		fmt.Printf("Error loading app configs: %s\n\n", err.Error())
		fmt.Println(" ----> Ending program.")
		return false
	}
	*configPtr = config
	fmt.Print("App config loaded with success!\n\n\n")
	return true
}

func loadSpotifyClient(config app.Config) client.SpotifyClient {
	spotifyConfig := config.Clients.Spotify
	spotifyClient := client.NewSpotifyClient(
		spotifyConfig.BaseUrl,
		spotifyConfig.AccountsUrl,
		spotifyConfig.ClientCredentials,
	)
	return spotifyClient
}

func runApp(spotifyClient client.SpotifyClient) bool {
	fmt.Println("Trying authenticate with client credentials...")
	oAuthResponse, err := spotifyClient.AuthenticateWithClientCredentials()
	if err != nil {
		fmt.Printf("Error authenticating: %s\n\n", err.Error())
		return false
	}
	fmt.Println("Authentication succeeded!")
	if body, err3 := json.Marshal(oAuthResponse); err3 == nil && body != nil {
		fmt.Println(string(body))
	}
	return true
}
