package feature

import (
	"github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/example/common/dto"
	"github.com/truc-engine/truc/util"
	"go.uber.org/zap"
)

func init() {
	engine.RegisterEndpoint(common.Engine, "/user/ping", Ping)
	engine.RegisterEventHandler("user.ping", func(payload *dto.UserPingRequest) {
		util.Logger.Info("User ping", zap.Any("payload", payload))
	})
}

type Ctx = engine.Context[dto.UserPingRequest, dto.UserPingResponse]

type CtxRes = engine.Res[dto.UserPingResponse]

func Ping(c *Ctx) *CtxRes {
	return c.Ok(&dto.UserPingResponse{Message: "User"})
}
