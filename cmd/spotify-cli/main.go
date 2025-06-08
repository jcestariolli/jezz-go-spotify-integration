package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/albums"
	"jezz-go-spotify-integration/internal/artists"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/config"
)

func main() {
	var spotifyConfig config.Config
	if !loadConfig(&spotifyConfig) {
		return
	}
	authService := loadAuthService(spotifyConfig)
	artistsSvc, albumSvc := loadServices(spotifyConfig, authService)

	runAppFixedCalls(*artistsSvc, *albumSvc)
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

func loadServices(cfg config.Config, authService *auth.Service) (*artists.Service, *albums.Service) {
	cliConfig := cfg.Client
	artistsSvc := loadArtistsService(cliConfig, authService)
	albumsSvc := loadAlbumsService(cliConfig, authService)
	return artistsSvc, albumsSvc
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

func runAppFixedCalls(artistsSvc artists.Service, albumsSvc albums.Service) {
	getArtist(artistsSvc, "7nzSoJISlVJsn7O0yTeMOB")
	getMultipleArtists(artistsSvc, "4DFhHyjvGYa9wxdHUjtDkc", "4lgrzShsg2FLA89UM2fdO5")
	getAlbum(albumsSvc, "1QJmLRcuIMMjZ49elafR3K")
	getAlbumForCountryMarket(albumsSvc, "1QJmLRcuIMMjZ49elafR3K")
	getMultipleAlbums(albumsSvc, "6JLTZPPzQDKjv6zkenbZnc", "0lw68yx3MhKflWFqCsGkIs")
	getMultipleAlbumsForCountryMarket(albumsSvc, "6JLTZPPzQDKjv6zkenbZnc", "0lw68yx3MhKflWFqCsGkIs")

}

func getArtist(svc artists.Service, artistId string) {
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

func getMultipleArtists(svc artists.Service, artistIds ...string) {
	fmt.Println("Trying to get multiple artists...")

	artistsResponse, err := svc.GetArtists(artistIds...)
	if err != nil {
		fmt.Println("✖ Getting multiple artists failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(artistsResponse); err3 == nil && body != nil {
		fmt.Println("✔ Artists obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple artists failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple artists failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getAlbum(svc albums.Service, albumId string) {
	fmt.Println("Trying to get an album...")

	albumResponse, err := svc.GetAlbum(nil, albumId)
	if err != nil {
		fmt.Println("✖ Getting album failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumResponse); err3 == nil && body != nil {
		fmt.Println("✔ Album obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting album failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting album failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getAlbumForCountryMarket(svc albums.Service, albumId string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get an album for " + countryMarketName + "'s market...")

	albumResponse, err := svc.GetAlbum(&countryMarketName, albumId)
	if err != nil {
		fmt.Println("✖ Getting album for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumResponse); err3 == nil && body != nil {
		fmt.Println("✔ Album obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting album failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting album for market failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getMultipleAlbums(svc albums.Service, albumIds ...string) {
	fmt.Println("Trying to get multiple albums...")

	albumsResponse, err := svc.GetAlbums(nil, albumIds...)
	if err != nil {
		fmt.Println("✖ Getting multiple albums failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumsResponse); err3 == nil && body != nil {
		fmt.Println("✔ Albums obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple albums failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple albums failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getMultipleAlbumsForCountryMarket(svc albums.Service, albumIds ...string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get multiple albums for " + countryMarketName + "'s market...")

	albumsResponse, err := svc.GetAlbums(&countryMarketName, albumIds...)
	if err != nil {
		fmt.Println("✖ Getting multiple albums for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumsResponse); err3 == nil && body != nil {
		fmt.Println("✔ Albums obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple albums failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple albums for market failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}
