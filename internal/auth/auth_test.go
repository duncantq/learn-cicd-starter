package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	empty_header := http.Header{}
	malformed_header := http.Header{}
	malformed_header.Add("Authorization", "invalid")

	tests := map[string]struct {
		header http.Header
	}{
		"empty header":     {header: empty_header},
		"malformed_header": {header: malformed_header},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := GetAPIKey(test.header)
			if err == nil {
				t.Fatalf("Expected err: got nil")
			}
		})
	}
}
