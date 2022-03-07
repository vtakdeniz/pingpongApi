package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestServerReturnsCorrectRespons(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		body         string
	}{
		{
			description:  "Server returns correct response",
			route:        "/ping",
			expectedCode: 200,
			body:         `{"pongs":["pong","pong","pong"]}`,
		},
	}

	app := initServer()

	for _, test := range tests {
		ping := Ping{Times: 3}
		pingJson, _ := json.Marshal(ping)
		req := httptest.NewRequest("POST", test.route, strings.NewReader(string(pingJson)))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, _ := app.Test(req, 1)
		body, _ := ioutil.ReadAll(resp.Body)
		bodyObj := Pong{}

		json.Unmarshal(body, &bodyObj)
		bodyJSON, _ := json.Marshal(bodyObj)
		bodyJSONString := string(bodyJSON)

		assert.JSONEq(t, test.body, bodyJSONString, test.description)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCraftResponse(t *testing.T) {
	tests := []struct {
		description string
		want        Pong
		times       int
	}{
		{
			description: "Function alters pong object correctly",
			want: Pong{[]string{
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
			},
			},
			times: 5,
		},
		{
			description: "Function alters pong object correctly",
			want: Pong{[]string{
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
				"pong",
			},
			},
			times: 13,
		},
		{
			description: "Function alters pong object correctly",
			want: Pong{[]string{
				"pong",
			},
			},
			times: 1,
		},
	}

	for _, test := range tests {
		pong := Pong{}
		CraftResponse(&pong, test.times)
		assert.Equalf(t, pong, test.want, test.description)
	}
}
