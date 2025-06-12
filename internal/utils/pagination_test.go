package utils

import (
	"jezz-go-spotify-integration/internal/model"
	"testing"

	"github.com/samber/lo"
)

func TestValidatePaginationParams(t *testing.T) {
	type args struct {
		limit  *model.Limit
		offset *model.Offset
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return nil error when all params are valid",
			args: args{
				limit:  lo.ToPtr(model.Limit(10)),
				offset: lo.ToPtr(model.Offset(10)),
			},
			wantErr: false,
		},
		{
			name: "should return nil error when all params are nil",
			args: args{
				limit:  nil,
				offset: nil,
			},
			wantErr: false,
		},
		{
			name: "should return error when limit is bellow 0",
			args: args{
				limit:  lo.ToPtr(model.Limit(-1)),
				offset: lo.ToPtr(model.Offset(10)),
			},
			wantErr: true,
		},
		{
			name: "should return error when limit is above 50",
			args: args{
				limit:  lo.ToPtr(model.Limit(51)),
				offset: lo.ToPtr(model.Offset(10)),
			},
			wantErr: true,
		},
		{
			name: "should return error when offset is bellow 0",
			args: args{
				limit:  lo.ToPtr(model.Limit(10)),
				offset: lo.ToPtr(model.Offset(-1)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePaginationParams(tt.args.limit, tt.args.offset); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePaginationParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
