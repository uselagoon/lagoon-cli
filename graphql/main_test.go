package graphql

import (
	"testing"

	"github.com/spf13/viper"
)

func TestValidateToken(t *testing.T) {
	got := hasValidToken("test")
	if got == true {
		t.Error("HasValidToken should fail as no token is set by default", got)
	}
	// set a valid token
	viper.Set(
		"lagoons.test.token",
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1NzI3Nzk5MjYsImV4cCI6NDc1OTk4OTUzMSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.9XKxJps00mGzgneaEp0nmI8aXlolMrD-Do2IooTP7d0",
	)
	got = hasValidToken("test")
	if got == false {
		t.Error("HasValidToken should not fail once a token is set", got)
	}
	got = VerifyTokenExpiry("test")
	if got == false {
		t.Error("ValidateToken should not fail if the token is valid", got)
	}
	// set an expired token
	viper.Set(
		"lagoons.test.token",
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1NDEyNDM0MTMsImV4cCI6MTU0MTI0NDYxNCwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.2wkAPeyVHAhGps_OEe_a8RXmv7_9GwP4ttFsjrLPZ84",
	)
	got = VerifyTokenExpiry("test")
	if got == true {
		t.Error("ValidateToken should fail if the token is invalid or expired", got)
	}
}
