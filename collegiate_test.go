package mwgoapi

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCollegiate_UnmarshalJSON(t *testing.T) {
	hello, _ := os.ReadFile("./responses/hello.json")

	tests := []struct {
		name    string
		in      []byte
		wantErr bool
	}{
		{
			name:    "happy path",
			in:      hello,
			wantErr: false,
		},
		{
			name:    "unhappy path",
			in:      []byte(`{"meta": "invalid"}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.Unmarshal(tt.in, &Collegiate{}); (err != nil) != tt.wantErr {
				t.Errorf("Collegiate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
