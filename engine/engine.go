package engine

import "github.com/nats-io/nats.go"

type Engine struct {
	Nats *nats.Conn
}

func NewEngine() *Engine {
	return &Engine{
		Nats: new(nats.Conn),
	}
}
