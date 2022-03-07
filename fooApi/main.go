package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Ping struct {
	Times int `json:"times"`
}

type Pong struct {
	Pongs []string `json:"pongs"`
}

func fetch(baseUrl string, pingObj Ping) (Pong, error) {
	ping, _ := json.Marshal(pingObj)
	res, err := http.Post(fmt.Sprintf("%s/ping", baseUrl), fiber.MIMEApplicationJSON, strings.NewReader(string(ping)))
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

func main() {
	pong, err := fetch("http://localhost:8080", Ping{Times: 12})
	if err != nil {
		log.Fatalf("Error fetching pongs")
	}
	fmt.Printf("%v", pong)
}
