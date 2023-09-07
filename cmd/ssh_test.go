package cmd

import (
	"testing"
)

func Test_generateSSHConnectionString(t *testing.T) {
	type args struct {
		project     string
		environment string
		lagoon      map[string]string
		service     string
		container   string
		isPortal    bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1 - no service or container",
			args: args{
				lagoon: map[string]string{
					"hostname": "lagoon.example.com",
					"port":     "22",
					"username": "example-com-main",
				},
			},
			want: `ssh -t -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -p 22 example-com-main@lagoon.example.com`,
		},
		{
			name: "test1 - service only, no container",
			args: args{
				lagoon: map[string]string{
					"hostname": "lagoon.example.com",
					"port":     "22",
					"username": "example-com-main",
				},
				service: "cli",
			},
			want: `ssh -t -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -p 22 example-com-main@lagoon.example.com service=cli`,
		},
		{
			name: "test3 - service and container",
			args: args{
				lagoon: map[string]string{
					"hostname": "lagoon.example.com",
					"port":     "22",
					"username": "example-com-main",
				},
				service:   "nginx-php",
				container: "php",
			},
			want: `ssh -t -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -p 22 example-com-main@lagoon.example.com service=nginx-php container=php`,
		},
		{
			name: "test4",
			args: args{
				lagoon: map[string]string{
					"hostname": "lagoon.example.com",
					"port":     "22",
					"username": "example-com-main",
					"sshKey":   "/home/user/.ssh/my-key",
				},
				service:   "cli",
				container: "cli",
			},
			want: `ssh -t -i /home/user/.ssh/my-key -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -p 22 example-com-main@lagoon.example.com service=cli container=cli`,
		},
		{
			name: "test5 - sshportal",
			args: args{
				lagoon: map[string]string{
					"hostname": "lagoon.example.com",
					"port":     "22",
					"username": "example-com-main",
					"sshKey":   "/home/user/.ssh/my-key",
				},
				isPortal:  true,
				service:   "cli",
				container: "cli",
			},
			want: `ssh -t -i /home/user/.ssh/my-key -p 22 example-com-main@lagoon.example.com service=cli container=cli`,
		},
		{
			name: "test6 - sshportal",
			args: args{
				lagoon: map[string]string{
					"hostname": "lagoon.example.com",
					"port":     "22",
					"username": "example-com-main",
				},
				isPortal:  true,
				service:   "cli",
				container: "cli",
			},
			want: `ssh -t -p 22 example-com-main@lagoon.example.com service=cli container=cli`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdProjectName = tt.args.project
			cmdProjectEnvironment = tt.args.environment

			if got := generateSSHConnectionString(tt.args.lagoon, tt.args.service, tt.args.container, tt.args.isPortal); got != tt.want {
				t.Errorf("generateSSHConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
