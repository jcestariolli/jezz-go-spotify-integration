package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/service"
)

func main() {
	var spotifyConfig config.Config
	if loadConfig(&spotifyConfig) != true {
		return
	}
	authService := loadServices(spotifyConfig)
	runApp(authService)
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

func loadServices(config config.Config) service.AuthService {
	fmt.Println("Loading spotify services...")
	cliConfig := config.Client
	authService := loadAuthService(cliConfig)
	fmt.Printf("✔ Services loaded! :)\n\n")
	return authService
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

func runApp(authService service.AuthService) {
	fmt.Println("Trying to authenticate application...")
	authSession, err := authService.AuthenticateApp()
	if err != nil {
		fmt.Println("✖ Authentication failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}
	fmt.Println("✔ Authentication succeeded! :)")
	if body, err3 := json.Marshal(authSession); err3 == nil && body != nil {
		fmt.Printf("╰┈➤%s\n", string(body))
	}
	fmt.Printf("\n")
	return
}
