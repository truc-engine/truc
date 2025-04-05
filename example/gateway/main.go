package main

import (
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/gateway"
	"github.com/truc-engine/truc/util"
)

func main() {
	util.Logger.Info("Gateway started on port " + common.Parmas.GatewayPort)
	gateway.StartServer(common.Engine, common.Parmas.GatewayPort)
}
