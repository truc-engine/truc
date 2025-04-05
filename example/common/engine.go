package common

import (
	"github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/util"
)

var Engine = engine.NewEngine()

type EngineParams struct {
	NatsUrl     string
	GatewayPort string
}

var Parmas = EngineParams{
	NatsUrl:     util.GetEnv("NATS_URL", "localhost:4222"),
	GatewayPort: util.GetEnv("GATEWAY_PORT", "9999"),
}

func init() {
	var err error
	Engine.Nats, err = engine.ConnectNats(Parmas.NatsUrl)
	if err != nil {
		panic(err)
	}
}
