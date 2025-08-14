package config

import "testing"

// TestIsFlagSet tests the IsFlagSet method on Lagoon Configs
func TestIsFlagSet(t *testing.T) {
	var testCases = map[string]struct {
		flags  []string
		flag   string
		expect bool
	}{
		"flag-set": {
			flags:  []string{"experimental", "verbose"},
			flag:   "experimental",
			expect: true,
		},
		"flag-not-set": {
			flags:  []string{"verbose"},
			flag:   "experimental",
			expect: false,
		},
		"flag-case-insensitive": {
			flags:  []string{"Experimental", "Verbose"},
			flag:   "experimental",
			expect: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			config := Config{Flags: tc.flags}
			result := config.IsFlagSet(tc.flag)
			if result != tc.expect {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
