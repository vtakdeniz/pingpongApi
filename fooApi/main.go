package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Ping struct {
	Times int `json:"times"`
}

type Pong struct {
	Pongs []string `json:"pongs"`
}

func fetch(baseUrl string, pingObj Ping) (Pong, error) {
	ping, _ := json.Marshal(pingObj)
	res, err := http.Post(fmt.Sprintf("%s/ping", baseUrl), "application/json", strings.NewReader(string(ping)))
	if err != nil {
		return Pong{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Pong{}, err
	}

	pong := new(Pong)
	err = json.Unmarshal(body, pong)
	if err != nil {
		return Pong{}, err
	}
	return *pong, nil
}

func FetchFromLocalHost(port int, pingObj Ping) error {
	ping, _ := json.Marshal(pingObj)
	res, err := http.Post(fmt.Sprintf("http://localhost:%d/ping", port), "application/json", strings.NewReader(string(ping)))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	pong := new(Pong)
	err = json.Unmarshal(body, pong)
	if err != nil {
		return err
	}

	fmt.Printf("%v", *pong)
	return nil
}

func main() {
	err := FetchFromLocalHost(8080, Ping{Times: 12})
	if err != nil {
		log.Fatalf("Error fetching pongs")
	}
}
