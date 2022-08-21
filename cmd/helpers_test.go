package cmd

import "testing"

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
