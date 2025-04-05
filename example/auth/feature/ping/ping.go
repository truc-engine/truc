package feature

import (
	"fmt"

	"github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/example/common"
	"github.com/truc-engine/truc/gateway"
)

type UserPingRequest struct {
	Message string `json:"message"`
}

type UserPingResponse struct {
	Message string `json:"message"`
}

func init() {
	gateway.RegisterEndpoint(common.Engine, "/auth/ping", Ping)
}

func Ping(c *Ctx) *CtxRes {
	c.Logger.Info("Ping " + c.Id)

	res, err := engine.SendRequest[UserPingRequest, UserPingResponse]("/user/pinzzg", &UserPingRequest{Message: c.Id})
	if err != nil {
		return c.ServerError()
	}
	c.Logger.Info(fmt.Sprintf("User ping response: %v", res.Ok))
	c.Logger.Info(fmt.Sprintf("User ping response: %v", res.Data.Message))

	return c.Ok(&AuthPingResponse{Message: c.Id})
}
