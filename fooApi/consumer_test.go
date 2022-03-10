//go:build consumer
// +build consumer

package main

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pact-foundation/pact-go/dsl"

	"github.com/pact-foundation/pact-go/types"
)

func InitPact() (pact *dsl.Pact, cleanUp func()) {
	pact = &dsl.Pact{
		Host:                     "localhost",
		Consumer:                 "foo",
		Provider:                 "bar",
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
		LogDir:                   "./pacts/logs",
	}
	cleanUp = func() { pact.Teardown() }
	return pact, cleanUp
}

func PublishPact() error {
	p := dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{"./pacts/foo-bar.json"},
		PactBroker:      "https://mediterranean.pactflow.io",
		BrokerToken:     "C05m5gQduXO-0fFGPYN6mw",
		ConsumerVersion: "1.0.0",
		Tags:            nil,
	})
	if err != nil {
		return err
	}
	return nil
}

func Test_ConsumerGetsCorrectAmountOfPongs(t *testing.T) {
	pact, cleanUp := InitPact()
	defer cleanUp()

	pingObj := Ping{Times: 5}
	res := Pong{Pongs: []string{"pong", "pong", "pong", "pong", "pong"}}

	pact.AddInteraction().Given("Ping Object").UponReceiving("Pong Object").WithRequest(
		dsl.Request{
			Method: "POST",
			Path:   dsl.String("/ping"),
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body: dsl.Like(pingObj),
		},
	).
		WillRespondWith(dsl.Response{
			Status: http.StatusOK,
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body: dsl.Like(res),
		})

	err := pact.Verify(func() error {
		return FetchFromLocalHost(pact.Server.Port, pingObj)
	})
	if err != nil {
		t.Fatal(err)
	}
	err = PublishPact()
	if err != nil {
		t.Fatal(err)
	}
}
