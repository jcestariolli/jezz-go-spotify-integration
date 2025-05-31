package main

import (
	_ "embed"
	"fmt"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/service"
)

//go:embed config/client_credentials.json
var clientCredentialsConfig []byte

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	fmt.Println("Loading credentials file....")
	clientCredentials, err := auth.LoadClientCredentialsFromFile(clientCredentialsConfig)
	if err != nil {
		fmt.Println("Error loading client_credentials file. Ending program.")
		return
	}
	fmt.Print("Credentials file loaded with success!\n\n\n")

	fmt.Println("Trying to authenticate to spotify....")
	spotifyAuthApiUrl := "https://accounts.spotify.com/api/token"
	spotifyService := service.NewSpotifyService(
		client.NewSpotifyClient(
			spotifyAuthApiUrl,
			clientCredentials,
		),
	)

	credentials, err := spotifyService.AuthenticateWithClientCredentials()
	if err != nil {
		fmt.Println("Error authenticating. Ending program.")
		return
	}
	fmt.Println("Authentication succeeded!")
	fmt.Printf("- Token: %s\n\n\n", credentials.Value)
}
