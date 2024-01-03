package custom

import (
	"reflect"
	"testing"
)

func TestConfig_readCommandfile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name        string
		description string
		args        args
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that reading a custom command file reads as expected",
			args: args{
				file: "test-data/deploytargets.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// c := &Commands{}
			c, err := readCommandfile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.readCommandfile() error = %v, wantErr %v", err, tt.wantErr)
			}
			seed, err := readCommandfile("test-data/deploytargets.yaml")
			if err != nil {
				t.Errorf("err %v", err)
			}
			if !reflect.DeepEqual(&seed, &c) {
				t.Errorf("Config.GetContext() = %v, want %v", &seed, c)
			}
		})
	}
}
