package engine

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func (e *Engine) PublishEvent(event string, payload proto.Message) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	e.Nats.Publish("event."+event, data)

	return nil
}

func (e *Engine) RegisterEventHandler(event string, handler func(payload any)) {
	event = "event." + event
	fmt.Println("Register event: " + event)
	e.Nats.QueueSubscribe(event, "event", func(msg *nats.Msg) {
		var c any
		err := json.Unmarshal(msg.Data, &c)
		if err != nil {
			fmt.Println(err)
			return
		}
		handler(c)
	})
}
