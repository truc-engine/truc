package engine

import (
	"encoding/json"
	"time"

	"github.com/truc-engine/truc/util"
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
