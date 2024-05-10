package client

import (
	"fmt"
	"testing"
)

func Test_resolvePureHost(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{in: "bank-de.api.local.ownera.io:80", out: "bank-de.api.local.ownera.io"},
		{in: "https://bank-de.api.local.ownera.io:80", out: "bank-de.api.local.ownera.io"},
		{in: "bank-de.api.local.ownera.io:443", out: "bank-de.api.local.ownera.io"},
		{in: "h2c://bank-de.api.local.ownera.io:443", out: "bank-de.api.local.ownera.io"},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			want, got := tc.out, resolvePureHost(tc.in)
			if got != want {
				t.Fatalf("invalid host: want  `%s`, got `%s`", want, got)
			}
		})
	}
}
