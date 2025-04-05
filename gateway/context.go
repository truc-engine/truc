package gateway

import "go.uber.org/zap"

type C struct {
	Url      string
	Request  any
	Response any
	Id       string
}

type Res struct {
	Data    any
	Ok      bool
	Code    string
	Message string
}

type Context[I, O any] struct {
	Url     string
	Request *I
	Id      string
	Logger  *zap.Logger
}

type ContextResponse[O any] struct {
	Data    *O
	Ok      bool
	Code    string
	Message string
}
