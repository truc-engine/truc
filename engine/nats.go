package engine

import (
	"strings"

	"github.com/nats-io/nats.go"
)

var Nats *nats.Conn

func ConnectNats(url string) (*nats.Conn, error) {
	if url == "" {
		url = nats.DefaultURL
	}
	nc, err := nats.Connect(url)
	if err != nil {
		return nc, err
	}
	Nats = nc
	return nc, nil
}

func ParseEndpoint(url string) string {
	url = strings.Replace(url, "/", ".", -1)

	return "api" + url
}
