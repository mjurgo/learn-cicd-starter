package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr error
	}{
		{
			name: "Valid auth header",
			headers: http.Header{
				"Authorization": {"ApiKey asdfghjkl"},
			},
			wantKey: "asdfghjkl",
			wantErr: nil,
		},
		{
			name:    "No auth header",
			headers: http.Header{},
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if err != tc.wantErr {
				t.Fatalf("expected error: %v; got: %v", tc.wantErr, err)
			}

			if !reflect.DeepEqual(tc.wantKey, got) {
				t.Fatalf("expected: %#v; got: %#v", tc.wantKey, got)
			}
		})
	}
}
