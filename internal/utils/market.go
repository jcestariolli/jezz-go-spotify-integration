package utils

import (
	"github.com/pariz/gountries"
	"github.com/samber/lo"
	"jezz-go-spotify-integration/internal/model"
)

func GetMarketByCountryName(countryName *string) (*model.AvailableMarket, error) {
	var market *model.AvailableMarket
	if countryName != nil {
		country, err := gountries.New().FindCountryByName(*countryName)
		if err != nil {
			return nil, err
		}
		market = lo.ToPtr(model.AvailableMarket(country.Alpha2))

	}
	return market, nil
}
