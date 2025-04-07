package truc

import (
	"strings"

	"github.com/nats-io/nats.go"
)

func (e *Engine) ConnectNats(url string) (*nats.Conn, error) {
	if url == "" {
		url = e.Params.Nats
	}
	nc, err := nats.Connect(url)
	if err != nil {
		return nc, err
	}
	e.Nats = nc
	return nc, nil
}

func ParseEndpoint(url string) string {
	url = strings.Replace(url, "/", ".", -1)

	return "api" + url
}

func ParseEvent(event string) string {
	return "event." + event
}
