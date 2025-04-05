package feature

import (
	"github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/example/common/dto"
	"github.com/truc-engine/truc/gateway"
	"go.uber.org/zap"
)

type Ctx = engine.Context[dto.AuthPingRequest, dto.AuthPingResponse]

type CtxRes = engine.Res[dto.AuthPingResponse]

func init() {
	gateway.RegisterEndpoint(common.Engine, "/auth/ping", Ping)
}

func Ping(c *Ctx) *CtxRes {
	c.Logger.Info("Ping")

	userPingRes, err := engine.SendRequest[dto.UserPingRequest, dto.UserPingResponse]("/user/ping", &dto.UserPingRequest{})
	if err != nil {
		return c.ServerError()
	}

	c.Logger.Info("User ping response", zap.Any("userPingRes", userPingRes))

	return c.Ok(&dto.AuthPingResponse{Message: c.Id})
}
