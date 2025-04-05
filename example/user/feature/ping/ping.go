package feature

import (
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/gateway"
)

func init() {
	gateway.RegisterEndpoint(common.Engine, "/user/ping", Ping)
}

func Ping(c *Ctx) *CtxRes {
	return c.Ok(&UserPingResponse{Message: "User"})
}
