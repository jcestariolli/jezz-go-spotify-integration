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
		return
	}
	spotifyClient := loadSpotifyClient(config)

	if runApp(spotifyClient) != true {
		return
	}
}

func loadConfig(configPtr *app.Config) bool {
	fmt.Println("Loading app configs...")
	config, err := app.Load()
	if err != nil {
		fmt.Println("Error loading app configs :(")
		fmt.Printf("----> %s\n\n", err.Error())
		return false
	}
	*configPtr = config
	fmt.Println("Configs loaded with success! :)")
	return true
}

func loadSpotifyClient(config app.Config) client.SpotifyClient {
	fmt.Println("Loading spotify client...")
	spotifyConfig := config.Clients.Spotify
	spotifyClient := client.NewSpotifyClient(
		spotifyConfig.BaseUrl,
		spotifyConfig.AccountsUrl,
		spotifyConfig.ClientCredentials,
	)
	fmt.Println("Client loaded! :)")
	return spotifyClient
}

func runApp(spotifyClient client.SpotifyClient) bool {
	fmt.Println("Trying to authenticate with client credentials...")
	oAuthResponse, err := spotifyClient.AuthenticateWithClientCredentials()
	if err != nil {
		fmt.Println("Authentication failed :(")
		fmt.Printf("----> %s\n\n", err.Error())
		return false
	}
	fmt.Println("Authentication succeeded! :)")
	if body, err3 := json.Marshal(oAuthResponse); err3 == nil && body != nil {
		fmt.Printf("----> %s\n\n", string(body))
	}
	return true
}
