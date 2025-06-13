package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/samber/lo"
)

type DummyConfig struct {
	DummyRequiredField        string  `json:"dummy_required_field" yaml:"dummy_required_field" validate:"required"`
	DummyURLField             string  `json:"dummy_url_field" yaml:"dummy_url_field" validate:"url"`
	DummySimpleField          int     `json:"dummy_simple_field" yaml:"dummy_simple_field"`
	DummyPointerField         *string `json:"dummy_pointer_field" yaml:"dummy_pointer_field"`
	DummyRequiredPointerField *string `json:"dummy_required_pointer_field" yaml:"dummy_required_pointer_field" validate:"required"`
}

func TestConfig_validate(t *testing.T) {
	type fields struct {
		Config DummyConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should return nil when config has all fields properly filled",
			fields: fields{
				Config: DummyConfig{
					DummyRequiredField:        "dummy-field",
					DummyURLField:             "http://dummy.url",
					DummySimpleField:          1,
					DummyPointerField:         lo.ToPtr("dummy-pointer-value"),
					DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
				},
			},
			wantErr: false,
		},
		{
			name: "should return nil when config is missing fields that are not required",
			fields: fields{
				Config: DummyConfig{
					DummyRequiredField:        "dummy-field",
					DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
				},
			},
			wantErr: false,
		},
		{
			name: "should return nil when config is missing pointer fields that are not required",
			fields: fields{
				Config: DummyConfig{
					DummyRequiredField:        "dummy-field",
					DummyURLField:             "http://dummy.url",
					DummySimpleField:          1,
					DummyPointerField:         nil,
					DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when config is empty but there are required fields",
			fields: fields{
				Config: DummyConfig{},
			},
			wantErr: true,
		},
		{
			name: "should return error when config has a field with url validation but its value is not an url",
			fields: fields{
				Config: DummyConfig{
					DummyRequiredField:        "dummy-field",
					DummyURLField:             "not-an-url",
					DummySimpleField:          1,
					DummyPointerField:         lo.ToPtr("dummy-pointer-value"),
					DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when config has a required field not filled",
			fields: fields{
				Config: DummyConfig{
					DummyRequiredField:        "",
					DummyURLField:             "http://dummy.url",
					DummySimpleField:          1,
					DummyPointerField:         lo.ToPtr("dummy-pointer-value"),
					DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when config has a required pointer field not filled",
			fields: fields{
				Config: DummyConfig{
					DummyRequiredField:        "dummy-field",
					DummyURLField:             "http://dummy.url",
					DummySimpleField:          1,
					DummyPointerField:         lo.ToPtr("dummy-pointer-value"),
					DummyRequiredPointerField: nil,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.fields.Config)
			if tt.wantErr && (err == nil) {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	type fields struct {
		configDataFile string
	}
	type testCase struct {
		name    string
		fields  fields
		want    DummyConfig
		wantErr bool
	}
	tests := []testCase{
		{
			name: "should load yaml config with all fields with success",
			fields: fields{
				configDataFile: "dummy-config-all-filled.yml",
			},
			want: DummyConfig{
				DummyRequiredField:        "dummy-field",
				DummyURLField:             "http://dummy.url",
				DummySimpleField:          1,
				DummyPointerField:         lo.ToPtr("dummy-pointer-value"),
				DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
			},
			wantErr: false,
		},
		{
			name: "should load yaml config with empty fields with success",
			fields: fields{
				configDataFile: "dummy-config-empty.yml",
			},
			want: DummyConfig{
				DummyRequiredField:        "",
				DummyURLField:             "",
				DummySimpleField:          0,
				DummyPointerField:         nil,
				DummyRequiredPointerField: nil,
			},
			wantErr: false,
		},
		{
			name: "should load json config with all fields with success",
			fields: fields{
				configDataFile: "dummy-config-all-filled.json",
			},
			want: DummyConfig{
				DummyRequiredField:        "dummy-field",
				DummyURLField:             "http://dummy.url",
				DummySimpleField:          1,
				DummyPointerField:         lo.ToPtr("dummy-pointer-value"),
				DummyRequiredPointerField: lo.ToPtr("dummy-required-pointer-value"),
			},
			wantErr: false,
		},
		{
			name: "should load json config with empty fields with success",
			fields: fields{
				configDataFile: "dummy-config-empty.json",
			},
			want: DummyConfig{
				DummyRequiredField:        "",
				DummyURLField:             "",
				DummySimpleField:          0,
				DummyPointerField:         nil,
				DummyRequiredPointerField: nil,
			},
		},
		{
			name: "should return error when config is from an invalid format",
			fields: fields{
				configDataFile: "dummy-config-invalid-format.txt",
			},
			want:    DummyConfig{},
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

			config := DummyConfig{}
			err := loadConfig(data, &config)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(config, tt.want) {
				t.Errorf("loadConfig() got = %v, want %v", config, tt.want)
			}
		})
	}
}
