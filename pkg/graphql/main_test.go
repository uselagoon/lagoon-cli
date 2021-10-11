package graphql

import (
	"testing"

	"github.com/uselagoon/lagoon-cli/internal/lagoon"
)

func TestValidateToken(t *testing.T) {
	lc := lagoon.Config{
		Lagoons: map[string]lagoon.Context{},
	}
	got := hasValidToken(&lc, "test")
	if got == true {
		t.Error("HasValidToken should fail as no token is set by default", got)
	}
	// set the context with a valid token
	lc.Lagoons["test"] = lagoon.Context{
		Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1NzI3Nzk5MjYsImV4cCI6NDc1OTk4OTUzMSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.9XKxJps00mGzgneaEp0nmI8aXlolMrD-Do2IooTP7d0",
	}
	got = hasValidToken(&lc, "test")
	if got == false {
		t.Error("HasValidToken should not fail once a token is set", got)
	}
	got = VerifyTokenExpiry(&lc, "test")
	if got == false {
		t.Error("ValidateToken should not fail if the token is valid", got)
	}
	// set the context with an expired token
	lc.Lagoons["test"] = lagoon.Context{
		Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1NDEyNDM0MTMsImV4cCI6MTU0MTI0NDYxNCwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.2wkAPeyVHAhGps_OEe_a8RXmv7_9GwP4ttFsjrLPZ84",
	}
	got = VerifyTokenExpiry(&lc, "test")
	if got == true {
		t.Error("ValidateToken should fail if the token is invalid or expired", got)
	}
}
