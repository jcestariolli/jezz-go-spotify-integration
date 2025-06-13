package config

import (
	"os"
	"reflect"
	"testing"
)

func TestAppConfigLoader_Load(t *testing.T) {
	type fields struct {
		configDataFile string
	}
	tests := []struct {
		name    string
		fields  fields
		want    AppConfig
		wantErr bool
	}{
		{
			name: "should load yaml app config with success",
			fields: fields{
				configDataFile: "app-config-all-filled.yml",
			},
			want: AppConfig{
				Client: CliConfig{
					BaseURL:     "http://dummy.url",
					AccountsURL: "http://dummy.url",
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when loading yaml app config that misses required fields",
			fields: fields{
				configDataFile: "app-config-missing-required-fields.yml",
			},
			want:    AppConfig{},
			wantErr: true,
		},
		{
			name: "should return error when loading yaml app config with malformed url fields",
			fields: fields{
				configDataFile: "app-config-url-malformed.yml",
			},
			want:    AppConfig{},
			wantErr: true,
		},
		{
			name: "should load json app config with success",
			fields: fields{
				configDataFile: "app-config-all-filled.json",
			},
			want: AppConfig{
				Client: CliConfig{
					BaseURL:     "http://dummy.url",
					AccountsURL: "http://dummy.url",
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when loading json app config that misses required fields",
			fields: fields{
				configDataFile: "app-config-missing-required-fields.json",
			},
			want:    AppConfig{},
			wantErr: true,
		},
		{
			name: "should return error when loading json app config with malformed url fields",
			fields: fields{
				configDataFile: "app-config-url-malformed.json",
			},
			want:    AppConfig{},
			wantErr: true,
		},
		{
			name: "should return error when app file is from an invalid format",
			fields: fields{
				configDataFile: "app-config-invalid-format.txt",
			},
			want:    AppConfig{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.configDataFile == "" {
				t.Fatalf("Config data file cannot be empty")
			}
			data, errRead := os.ReadFile(testDataDir + "/" + tt.fields.configDataFile)
			if errRead != nil {
				t.Fatalf("Could not load config data file")
				return
			}

			a := AppConfigLoader{}
			got, err := a.Load(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}
