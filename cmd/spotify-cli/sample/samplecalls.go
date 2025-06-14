package sample

import (
	"encoding/json"
	"fmt"
	"jezz-go-spotify-integration/internal/service"
	"strings"
)

const (
	DefaultAlbumGroup     = "album"
	SingleAlbumGroup      = "single"
	AppearsOnAlgumGroup   = "appears_on"
	CompilationAlbumGroup = "compilation"
)

func RunAppSampleCalls(artistsSvc service.ArtistsService, albumsSvc service.AlbumsService, tracksSvc service.TracksService) {

	getArtist(artistsSvc, "7nzSoJISlVJsn7O0yTeMOB")
	getMultipleArtists(artistsSvc, "4DFhHyjvGYa9wxdHUjtDkc", "4lgrzShsg2FLA89UM2fdO5")

	getArtistAlbums(artistsSvc, "0k17h0D3J5VfsdmQ1iZtE9")
	getArtistAlbumsType(artistsSvc, "0k17h0D3J5VfsdmQ1iZtE9", DefaultAlbumGroup)
	getArtistAlbumsType(artistsSvc, "0k17h0D3J5VfsdmQ1iZtE9", SingleAlbumGroup, CompilationAlbumGroup)
	getArtistAlbumsType(artistsSvc, "0k17h0D3J5VfsdmQ1iZtE9", AppearsOnAlgumGroup)

	getArtistTopTracks(artistsSvc, "5LfGQac0EIXyAN8aUwmNAQ")

	getAlbum(albumsSvc, "1QJmLRcuIMMjZ49elafR3K")
	getAlbumForCountryMarket(albumsSvc, "4R3tXoorBpHji6Jdms8a4Q")

	getMultipleAlbums(albumsSvc, "4jvurVXLanQyP1rPZjbSln", "0lw68yx3MhKflWFqCsGkIs")
	getMultipleAlbumsForCountryMarket(albumsSvc, "6JLTZPPzQDKjv6zkenbZnc", "4M7bISEIiCfNN8EuLu8wc6")

	getAlbumTracks(albumsSvc, "1QJmLRcuIMMjZ49elafR3K")
	getAlbumTracksForCountryMarket(albumsSvc, "4R3tXoorBpHji6Jdms8a4Q")

	getNewReleases(albumsSvc)

	getTrack(tracksSvc, "3O5JIwSON3KBaoyMUsjLjn")
	getTrackForCountryMarket(tracksSvc, "4h6G18XTQMtNpwYIXnrZI6")

	getMultipleTracks(tracksSvc, "2C6h8jV6NzbS9o3JNQ6j7p", "3GylBJWB3nHyFjgEm62pMD")
	getMultipleTracksForCountryMarket(tracksSvc, "4VQu1ooCteGDynSZYUgvT4", "3Zjdqz7eOox8XU0zTCPL4P")

}

func getArtist(svc service.ArtistsService, artistID string) {
	fmt.Println("Trying to get an artist...")

	artistResponse, err := svc.GetArtist(artistID)
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

func getMultipleArtists(svc service.ArtistsService, artistIDs ...string) {
	fmt.Println("Trying to get multiple artists...")

	artistsResponse, err := svc.GetArtists(artistIDs...)
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

func getArtistAlbums(svc service.ArtistsService, artistID string) {
	fmt.Println("Trying to get all artist's album types ...")

	artistResponse, err := svc.GetArtistAlbums(nil, nil, nil, nil, artistID)
	if err != nil {
		fmt.Println("✖ Getting all artist's album types failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(artistResponse); err3 == nil && body != nil {
		fmt.Println("✔ All Artist's album types obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting all artist's album types failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting all artist's album types failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getArtistAlbumsType(svc service.ArtistsService, artistID string, albumTypes ...string) {
	albumTypesStr := strings.Join(albumTypes, " and ")
	fmt.Println("Trying to get artist's " + albumTypesStr + "s ...")

	artistResponse, err := svc.GetArtistAlbums(nil, &albumTypes, nil, nil, artistID)
	if err != nil {
		fmt.Println("✖ Getting artist's " + albumTypesStr + "s failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(artistResponse); err3 == nil && body != nil {
		fmt.Println("✔ Artist's " + albumTypesStr + "s obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting artist's " + albumTypesStr + "s failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting artist's " + albumTypesStr + "s failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getArtistTopTracks(svc service.ArtistsService, artistID string) {
	fmt.Println("Trying to get artist's top-tracks...")

	artistResponse, err := svc.GetArtistTopTracks(nil, artistID)
	if err != nil {
		fmt.Println("✖ Getting artist's top-tracks failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(artistResponse); err3 == nil && body != nil {
		fmt.Println("✔ Artist's top-tracks obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting artist's top-tracks failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting artist's top-tracks failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getAlbum(svc service.AlbumsService, albumID string) {
	fmt.Println("Trying to get an album...")

	albumResponse, err := svc.GetAlbum(nil, albumID)
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

func getAlbumForCountryMarket(svc service.AlbumsService, albumID string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get an album for " + countryMarketName + "'s market...")

	albumResponse, err := svc.GetAlbum(&countryMarketName, albumID)
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

func getMultipleAlbums(svc service.AlbumsService, albumIDs ...string) {
	fmt.Println("Trying to get multiple albums...")

	albumsResponse, err := svc.GetAlbums(nil, albumIDs...)
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

func getMultipleAlbumsForCountryMarket(svc service.AlbumsService, albumIDs ...string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get multiple albums for " + countryMarketName + "'s market...")

	albumsResponse, err := svc.GetAlbums(&countryMarketName, albumIDs...)
	if err != nil {
		fmt.Println("✖ Getting multiple albums for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumsResponse); err3 == nil && body != nil {
		fmt.Println("✔ Albums for market obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple albums for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple albums for market failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getAlbumTracks(svc service.AlbumsService, albumID string) {
	fmt.Println("Trying to get album's tracks...")

	albumResponse, err := svc.GetAlbumTracks(nil, nil, nil, albumID)
	if err != nil {
		fmt.Println("✖ Getting album's tracks failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumResponse); err3 == nil && body != nil {
		fmt.Println("✔ Album's tracks obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting album's tracks failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting album's tracks failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getAlbumTracksForCountryMarket(svc service.AlbumsService, albumID string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get an album's tracks for " + countryMarketName + "'s market...")

	albumResponse, err := svc.GetAlbumTracks(&countryMarketName, nil, nil, albumID)
	if err != nil {
		fmt.Println("✖ Getting album's tracks for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumResponse); err3 == nil && body != nil {
		fmt.Println("✔ Album's tracks for market obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting album's tracks for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting album's tracks for market failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getNewReleases(svc service.AlbumsService) {
	fmt.Println("Trying to get new releases...")

	albumResponse, err := svc.GetNewReleases(nil, nil)
	if err != nil {
		fmt.Println("✖ Getting new releases failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(albumResponse); err3 == nil && body != nil {
		fmt.Println("✔ New releases obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting new releases failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting new releases failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getTrack(svc service.TracksService, trackID string) {
	fmt.Println("Trying to get an track...")

	trackResponse, err := svc.GetTrack(nil, trackID)
	if err != nil {
		fmt.Println("✖ Getting track failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(trackResponse); err3 == nil && body != nil {
		fmt.Println("✔ Track obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting track failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting track failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getTrackForCountryMarket(svc service.TracksService, trackID string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get an track for " + countryMarketName + "'s market...")

	trackResponse, err := svc.GetTrack(&countryMarketName, trackID)
	if err != nil {
		fmt.Println("✖ Getting track for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(trackResponse); err3 == nil && body != nil {
		fmt.Println("✔ Track for market obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting track for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting track for market failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getMultipleTracks(svc service.TracksService, trackIDs ...string) {
	fmt.Println("Trying to get multiple tracks...")

	tracksResponse, err := svc.GetTracks(nil, trackIDs...)
	if err != nil {
		fmt.Println("✖ Getting multiple tracks failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(tracksResponse); err3 == nil && body != nil {
		fmt.Println("✔ Tracks obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple tracks failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple tracks failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}

func getMultipleTracksForCountryMarket(svc service.TracksService, trackIDs ...string) {
	countryMarketName := "Brazil"
	fmt.Println("Trying to get multiple tracks for " + countryMarketName + "'s market...")

	tracksResponse, err := svc.GetTracks(&countryMarketName, trackIDs...)
	if err != nil {
		fmt.Println("✖ Getting multiple tracks for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err.Error())
		return
	}

	if body, err3 := json.Marshal(tracksResponse); err3 == nil && body != nil {
		fmt.Println("✔ Tracks for market obtained! :)")
		fmt.Printf("╰┈➤%s\n\n", string(body))
		return
	} else if err3 != nil {
		fmt.Println("✖ Getting multiple tracks for market failed :(")
		fmt.Printf("╰┈➤%s\n\n", err3.Error())
		return
	}
	fmt.Println("✖ Getting multiple tracks for market failed :(")
	fmt.Printf("╰┈➤Body is empty\n\n")
}
