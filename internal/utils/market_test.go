package utils

import (
	"jezz-go-spotify-integration/internal/model"
	"reflect"
	"testing"

	"github.com/samber/lo"
)

func TestGetMarketByCountryName(t *testing.T) {
	type args struct {
		countryName *string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.AvailableMarket
		wantErr bool
	}{
		{
			name:    "should return nil when country name is nil",
			args:    args{countryName: nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "should return available market with alpha2 from country when country name is valid",
			args:    args{countryName: lo.ToPtr("Brazil")},
			want:    lo.ToPtr(model.AvailableMarket("BR")),
			wantErr: false,
		},
		{
			name:    "should return error when country name is invalid",
			args:    args{countryName: lo.ToPtr("Not Valid")},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMarketByCountryName(tt.args.countryName)
			if tt.wantErr && err == nil {
				t.Errorf("GetMarketByCountryName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarketByCountryName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
