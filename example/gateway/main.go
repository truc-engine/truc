package main

import (
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/gateway"
)

func main() {
	gateway.StartServer(common.Engine, common.Parmas.GatewayPort)
}
