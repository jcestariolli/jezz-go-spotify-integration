package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/config"
	"jezz-go-spotify-integration/internal/service"
)

func main() {
	var spotifyConfig config.Config
	if !loadConfig(&spotifyConfig) {
		return
	}
	catalogService := loadServices(spotifyConfig)
	if catalogService == nil {
		return
	}
	runApp(*catalogService)
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

func loadServices(config config.Config) *service.CatalogService {
	fmt.Println("Loading catalog service...")
	cliConfig := config.Client
	catalogService, err := loadCatalogService(cliConfig)
	if err != nil {
		fmt.Println("✖ Catalog service loading failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return nil
	}
	fmt.Printf("✔ Cataçpg service loaded! :)\n\n")
	return catalogService
}

func loadCatalogService(cliConfig config.CliConfig) (*service.CatalogService, error) {
	return service.NewCatalogService(
		cliConfig.BaseUrl,
		cliConfig.AccountsUrl,
		cliConfig.CliCredentials,
	)
}

func runApp(catalogService service.CatalogService) {
	getArtist(catalogService, "7nzSoJISlVJsn7O0yTeMOB")
	getMultipleArtists(catalogService, "4DFhHyjvGYa9wxdHUjtDkc", "4lgrzShsg2FLA89UM2fdO5")
}

func getArtist(catalogService service.CatalogService, artistId string) bool {
	fmt.Println("Trying to get an artist...")
	artist, err := catalogService.GetArtist(artistId)

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

func getMultipleArtists(catalogService service.CatalogService, artistIds ...string) bool {
	fmt.Println("Trying to get multiple artists...")
	artists, err := catalogService.GetArtists(artistIds...)

	if err != nil {
		fmt.Println("✖ Getting multiple artists failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return true
	}
	fmt.Println("✔ Artists obtained! :)")
	if body, err3 := json.Marshal(artists); err3 == nil && body != nil {
		fmt.Printf("╰┈➤%s\n", string(body))
	}
	fmt.Printf("\n")
	return false
}
