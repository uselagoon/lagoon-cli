package cmd

import (
	"reflect"
	"testing"

	"github.com/guregu/null"
	"github.com/spf13/pflag"
)

func Test_makeSafe(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "slash in name",
			in:   "Feature/Branch",
			want: "feature-branch",
		},
		{
			name: "noslash in name",
			in:   "Feature-Branch",
			want: "feature-branch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSafe(tt.in); got != tt.want {
				t.Errorf("makeSafe() go %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "generate hash",
			in:   "feature-branch",
			want: "011122006d017c21d1376add9f7f65b43555a455",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashString(tt.in); got != tt.want {
				t.Errorf("hashString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortenEnvironment(t *testing.T) {
	type args struct {
		project     string
		environment string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "really long environment name with slash and capitals",
			args: args{
				environment: makeSafe("Feature/Really-Exceedingly-Long-Environment-Name-For-A-Branch"),
				project:     "this-is-my-project",
			},
			want: "feature-really-exceedingly-long-env-dc8c",
		},
		{
			name: "short environment name",
			args: args{
				environment: makeSafe("Feature/Branch"),
				project:     "this-is-my-project",
			},
			want: "feature-branch",
		},
		{
			name: "short environment name",
			args: args{
				environment: makeSafe("release/1.2.3"),
				project:     "this-is-my-project",
			},
			want: "release-1-2-3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortenEnvironment(tt.args.project, tt.args.environment); got != tt.want {
				t.Errorf("shortenEnvironment() got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flagStringNullValueOrNil(t *testing.T) {
	type args struct {
		flags map[string]interface{}
		flag  string
	}
	tests := []struct {
		name    string
		args    args
		want    *null.String
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				flags: map[string]interface{}{
					"build-image": nil,
				},
				flag: "build-image",
			},
			want: nil,
		},
		{
			name: "test1",
			args: args{
				flags: map[string]interface{}{
					"build-image": "",
				},
				flag: "build-image",
			},
			want: &null.String{},
		},
		{
			name: "test1",
			args: args{
				flags: map[string]interface{}{
					"build-image": "buildimage",
				},
				flag: "build-image",
			},
			want: func() *null.String {
				l := null.StringFrom("buildimage")
				return &l
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := &pflag.FlagSet{}
			for k, v := range tt.args.flags {
				flags.StringP(k, "", "", "")
				if v != nil {
					flags.Set(k, v.(string))
				}
			}
			got, err := flagStringNullValueOrNil(flags, tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("flagStringNullValueOrNil() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flagStringNullValueOrNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitInvokeTaskArguments(t *testing.T) {
	type args struct {
		invokedTaskArguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Standard parsing, single argument",
			args: args{
				invokedTaskArguments: []string{
					"KEY1=VALUE1",
				},
			},
			want: map[string]string{
				"KEY1": "VALUE1",
			},
			wantErr: false,
		},
		{
			name: "Invalid Input, multiple arguments",
			args: args{
				invokedTaskArguments: []string{
					"KEY1=VALUE1",
					"INVALID_ARGUMENT",
				},
			},
			want:    map[string]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitInvokeTaskArguments(tt.args.invokedTaskArguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitInvokeTaskArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitInvokeTaskArguments() got = %v, want %v", got, tt.want)
			}
		})
	}
}
