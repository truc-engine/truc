package feature

import (
	"github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/example/common/dto"
	"github.com/truc-engine/truc/gateway"
)

func init() {
	gateway.RegisterEndpoint(common.Engine, "/user/ping", Ping)
}

type Ctx = engine.Context[dto.UserPingRequest, dto.UserPingResponse]

type CtxRes = engine.Res[dto.UserPingResponse]

func Ping(c *Ctx) *CtxRes {
	return c.Ok(&dto.UserPingResponse{Message: "User"})
}
