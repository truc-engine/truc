package engine

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func PublishEvent(event string, payload proto.Message) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	Nats.Publish("event."+event, data)

	return nil
}

func RegisterEventHandler[I any](event string, handler func(payload *I)) {
	event = "event." + event
	fmt.Println("Register event: " + event)
	Nats.QueueSubscribe(event, "event", func(msg *nats.Msg) {
		var c I
		err := json.Unmarshal(msg.Data, &c)
		if err != nil {
			fmt.Println(err)
			return
		}
		handler(&c)
	})
}
