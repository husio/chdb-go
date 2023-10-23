package chdb

import (
	"errors"
	"testing"
)

func TestChdb(t *testing.T) {
	cases := map[string]struct {
		Query   string
		Want    string
		Path    string
		Format  string
		WantErr error
	}{
		"select const 1": {
			Query:  "SELECT 1",
			Want:   "1\n",
			Format: "CSV",
		},
		"select const string": {
			Query:  `SELECT 'hello'`,
			Want:   "\"hello\"\n",
			Format: "CSV",
		},
		"select version": {
			Query:  `SELECT version()`,
			Want:   "\"23.6.1.1\"\n",
			Format: "CSV",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := Query(tc.Query, tc.Format, tc.Path)
			if !errors.Is(err, tc.WantErr) {
				t.Fatalf("Unexpected error: %+v", err)
			}
			defer result.Close()
			if tc.WantErr != nil {
				return
			}

			got := result.Bytes()
			if got == nil {
				t.Fatal("nil result")
			}
			if got := string(got); tc.Want != got {
				t.Logf("want %q", tc.Want)
				t.Logf(" got %q", got)
				t.Fatal("Unexpected result")
			}
		})
	}
}
