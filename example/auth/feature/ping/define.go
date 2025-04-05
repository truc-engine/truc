package feature

import (
	"github.com/truc-engine/truc/engine"
)

type AuthPingRequest struct {
	Message string `json:"message"`
}

type AuthPingResponse struct {
	Message string `json:"message"`
}

type Ctx = engine.Context[AuthPingRequest, AuthPingResponse]

type CtxRes = engine.Res[AuthPingResponse]
