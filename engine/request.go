package engine

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/truc-engine/truc/util"
	"go.uber.org/zap"
)

func SendRequest[I, O any](url string, payload *I) (*Res[O], error) {
	url = ParseEndpoint(url)

	context := C{
		Url:     url,
		Request: payload,
		Id:      util.NewUuidV7(),
	}
	data, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}
	res, err := Nats.Request(url, data, 600*time.Second)
	if err != nil {
		return nil, err
	}
	c := new(Res[O])
	err = json.Unmarshal(res.Data, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

type Handler[I, O any] func(c *Context[I, O]) *Res[O]

func RegisterEndpoint[I, O any](e *Engine, url string, handler Handler[I, O]) {
	url = ParseEndpoint(url)
	fmt.Println("Register endpoint: " + url)

	_, err := e.Nats.QueueSubscribe(url, "API", func(msg *nats.Msg) {
		c := new(Context[I, O])
		err := json.Unmarshal(msg.Data, &c)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Logger = util.Logger.With(zap.String("id", c.Id))

		res := handler(c)
		resData, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		e.Nats.Publish(msg.Reply, resData)
	})

	if err != nil {
		panic(err)
	}
}
