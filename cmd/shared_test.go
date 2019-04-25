package cmd

import (
	"github.com/spf13/viper"
	"testing"
)

func TestValidateToken(t *testing.T) {
	got := ValidateToken()
	if got == true {
		t.Error("No token is set by default", got)
	}
	viper.Set("lagoon_token", "testtoken")
	got = ValidateToken()
	if got == false {
		t.Error("ValidateToken should not fail once a token is set", got)
	}
}
