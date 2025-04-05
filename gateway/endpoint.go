package gateway

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/util"
	"go.uber.org/zap"
)

type Handler[I, O any] func(c *engine.Context[I, O]) *engine.Res[O]

func RegisterEndpoint[I, O any](e *engine.Engine, url string, handler Handler[I, O]) {
	url = engine.ParseEndpoint(url)
	fmt.Println("Register endpoint: " + url)

	_, err := e.Nats.QueueSubscribe(url, "API", func(msg *nats.Msg) {
		c := new(engine.Context[I, O])
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
