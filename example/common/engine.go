package common

import (
	"fmt"

	"github.com/truc-engine/truc/engine"
)

var Engine = engine.NewEngine()

func init() {
	var err error
	Engine.Nats, err = engine.ConnectNats("")
	if err != nil {
		panic(err)
	}
	fmt.Println("Nats connected")
}
