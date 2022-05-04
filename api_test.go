package mwgoapi_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/murat/mwgoapi"
)

func TestNewClient(t *testing.T) {
	type args struct {
		client *http.Client
		url    string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "use collegiate api url",
			args: args{
				client: &http.Client{},
				url:    "",
				key:    "",
			},
		},
		{
			name: "use thesaurus api url",
			args: args{
				client: &http.Client{},
				url:    "https://www.dictionaryapi.com/api/v3/references/thesaurus/json",
				key:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := mwgoapi.NewClient(tt.args.client, tt.args.url, tt.args.key)
			if tt.args.url != "" && h.BaseURL != tt.args.url {
				t.Errorf("h.BaseURL expected(%s), got(%s)", mwgoapi.BaseURL, h.BaseURL)
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	helloResponse, err := os.ReadFile("./responses/collegiate.json")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		word    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "success",
			word:    "hello",
			want:    helloResponse,
			wantErr: false,
		},
		{
			name:    "fail",
			word:    "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Add("content-type", "application/json")
				if tt.wantErr {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(tt.want)
			}))
			defer srv.Close()

			c := mwgoapi.NewClient(srv.Client(), srv.URL, "xxx")
			got, err := c.Get(tt.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.Get() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
