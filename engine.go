package engine

import (
	"database/sql"

	"github.com/nats-io/nats.go"
)

type Engine struct {
	Nats   *nats.Conn
	Params EngineParams
	Db     *sql.DB
}

type EngineParams struct {
	Nats string
	Pg   string
}

func NewEngine(params EngineParams) *Engine {
	return &Engine{
		Nats:   new(nats.Conn),
		Params: params,
	}
}
