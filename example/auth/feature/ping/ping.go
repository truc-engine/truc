package feature

import (
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/gateway"
	"github.com/truc-engine/truc/util"
)

func init() {
	gateway.RegisterEndpoint(common.Engine, "/auth/ping", Ping)
}

func Ping(c *Ctx) *Ctx {
	id := util.NewUuidV7()
	c.Logger.Info("Ping " + id)
	c.Response = &PingResponse{Message: id}
	return c
}
