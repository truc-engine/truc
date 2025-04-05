package engine

import (
	"go.uber.org/zap"
)

type C struct {
	Url      string
	Request  any
	Response any
	Id       string
}

type Context[I, O any] struct {
	Url     string
	Request *I
	Id      string
	Logger  *zap.Logger
}

type Res[O any] struct {
	Data       *O     `json:"data"`
	Ok         bool   `json:"ok"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int
}

func (c *Context[I, O]) Ok(data *O) *Res[O] {
	res := &Res[O]{
		Data:       data,
		Ok:         true,
		Code:       "200",
		Message:    "Success",
		StatusCode: 200,
	}
	return res
}

func (c *Context[I, O]) Error(code string, message string) *Res[O] {
	res := &Res[O]{
		Data:       nil,
		Ok:         false,
		Code:       code,
		Message:    message,
		StatusCode: 400,
	}
	return res
}

func (c *Context[I, O]) ServerError(message ...string) *Res[O] {
	msg := "Internal Server Error"
	if len(message) > 0 {
		msg = message[0]
	}
	res := &Res[O]{
		Data:       nil,
		Ok:         false,
		Code:       "SERVER_ERROR",
		Message:    msg,
		StatusCode: 500,
	}
	return res
}
