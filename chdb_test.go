package chdb

import (
	"strings"
	"testing"
)

var testCases = []struct {
	query       string
	expected    string
	description string
}{
	{"SELECT 1", "1", "Basic Select"},
	{"SELECT 'hello'", "\"hello\"", "Basic Select"},
}

func TestChdb(t *testing.T) {
	for _, test := range testCases {
		observed, err := Query(test.query, "CSV")
		if err != nil {
			t.Fatal(err)
		}
		observed = strings.Replace(observed, "\n", "", -1)
		if observed != test.expected {
			t.Fatalf("%s: %s is not %s", test.query, observed, test.expected)
		}
	}
}
