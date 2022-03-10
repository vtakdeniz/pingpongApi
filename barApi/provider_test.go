//go:build provider
// +build provider

package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

type Settings struct {
	Host                       string
	ConsumerName               string
	ProviderName               string
	PactURL                    string
	PublishVerificationResults bool
	FailIfNoPactsFound         bool
	DisableToolValidityCheck   bool
	BrokerBaseURL              string
	BrokerToken                string
	ProviderVersion            string
	PactFileWriteMode          string
}

func (s *Settings) InitSettings() {
	s.Host = "127.0.0.1"
	s.ConsumerName = "foo"
	s.ProviderName = "bar"
	s.PactURL = "https://mediterranean.pactflow.io/pacts/provider/bar/consumer/foo/latest"
	s.PublishVerificationResults = true
	s.FailIfNoPactsFound = true
	s.DisableToolValidityCheck = true
	s.BrokerBaseURL = "https://mediterranean.pactflow.io"
	s.BrokerToken = "C05m5gQduXO-0fFGPYN6mw"
	s.ProviderVersion = "1.0.0"
	s.PactFileWriteMode = "merge"
}

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()

	go StartServer(port)

	settings := Settings{}
	settings.InitSettings()

	pact := dsl.Pact{
		Consumer:                 settings.ConsumerName,
		Provider:                 settings.ProviderName,
		Host:                     settings.Host,
		DisableToolValidityCheck: settings.DisableToolValidityCheck,
	}
	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://%s:%d", settings.Host, port),
		PactURLs:                   []string{settings.PactURL},
		BrokerURL:                  settings.BrokerBaseURL,
		BrokerToken:                settings.BrokerToken,
		FailIfNoPactsFound:         settings.FailIfNoPactsFound,
		PublishVerificationResults: settings.PublishVerificationResults,
		ProviderVersion:            settings.ProviderVersion,
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

}
