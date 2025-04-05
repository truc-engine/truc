package feature

import "github.com/truc-engine/truc/engine"

type UserPingRequest struct {
	Message string `json:"message"`
}

type UserPingResponse struct {
	Message string `json:"message"`
}

type Ctx = engine.Context[UserPingRequest, UserPingResponse]

type CtxRes = engine.Res[UserPingResponse]
