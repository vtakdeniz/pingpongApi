package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fetch(t *testing.T) {
	type args struct {
		pingObj Ping
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Fetches wanted data correctly - 3 pongs",
			args: args{
				pingObj: Ping{
					Times: 3,
				},
			},
			want:    `{"pongs":["pong","pong","pong"]}`,
			wantErr: false,
		},
		{
			name: "Fetches wanted data correctly - 24 pongs",
			args: args{
				pingObj: Ping{
					Times: 24,
				},
			},
			want:    `{"pongs":["pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong","pong"]}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, tt.want)
			}))
			defer svr.Close()
			log.Printf("%v", tt.args.pingObj)
			pong, err := fetch(svr.URL, tt.args.pingObj)
			if err != nil {
				t.Errorf("expected err to be nil got %v", err)
			}
			pongbyte, err := json.Marshal(pong)
			if err != nil {
				t.Errorf("expected err to be nil got %v", err)
			}
			assert.JSONEq(t, string(pongbyte), tt.want, tt.name)
		})
	}
}
