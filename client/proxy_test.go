package client

import "testing"

func Test_resolvePureHost(t *testing.T) {
	got := resolvePureHost("bank-de.api.local.ownera.io:80")
	want := "bank-de.api.local.ownera.io"
	if got != want {
		t.Fatalf("invalid host: want `%s`, got `%s`", want, got)
	}
}
