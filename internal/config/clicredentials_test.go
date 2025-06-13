package config

import (
	"os"
	"reflect"
	"testing"
)

func TestCliCredentialsConfigLoader_Load(t *testing.T) {
	type fields struct {
		configDataFile string
	}
	tests := []struct {
		name    string
		fields  fields
		want    CliCredentials
		wantErr bool
	}{
		{
			name: "should load yaml client credentials config with success",
			fields: fields{
				configDataFile: "cli-credentials-config-all-filled.yml",
			},
			want: CliCredentials{
				Id:     "DUMMY_CLIENT_ID",
				Secret: "DUMMY_CLIENT_SECRET",
			},
			wantErr: false,
		},
		{
			name: "should return error when loading yaml client credentials config that misses required fields",
			fields: fields{
				configDataFile: "cli-credentials-config-missing-required-fields.yml",
			},
			want:    CliCredentials{},
			wantErr: true,
		},
		{
			name: "should load json client credentials config with success",
			fields: fields{
				configDataFile: "cli-credentials-config-all-filled.json",
			},
			want: CliCredentials{
				Id:     "DUMMY_CLIENT_ID",
				Secret: "DUMMY_CLIENT_SECRET",
			},
			wantErr: false,
		},
		{
			name: "should return error when loading json client credentials config that misses required fields",
			fields: fields{
				configDataFile: "cli-credentials-config-missing-required-fields.json",
			},
			want:    CliCredentials{},
			wantErr: true,
		},
		{
			name: "should return error when client credentials file is from an invalid format",
			fields: fields{
				configDataFile: "cli-credentials-config-invalid-format.txt",
			},
			want:    CliCredentials{},
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

			a := CliCredentialsConfigLoader{}
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
