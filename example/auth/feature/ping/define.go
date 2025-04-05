package feature

import "github.com/truc-engine/truc/gateway"

type PingRequest struct {
	Message string `json:"message"`
}

type PingResponse struct {
	Message string `json:"message"`
}

type Ctx = gateway.Context[PingRequest, PingResponse]
