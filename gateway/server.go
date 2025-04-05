package gateway

import (
	"encoding/json"
	"net/http"
	"time"

	e "github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/util"
)

type HttpResponse struct {
	Data    *any   `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func StartServer(engine *e.Engine, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := e.ParseEndpoint(r.URL.String())
		var payload any

		// Decode the JSON payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		context := e.C{
			Url:     url,
			Request: payload,
			Id:      util.NewUuidV7(),
		}
		data, err := json.Marshal(context)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		natResponse, err := engine.Nats.Request(url, data, 600*time.Second)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		c := new(e.Res[any])
		err = json.Unmarshal(natResponse.Data, &c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if c.StatusCode == 0 {
			c.StatusCode = 500
			c.Code = "SERVER_ERROR"
			c.Data = nil
			c.Message = "Internal Server Error"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(c.StatusCode)
		json.NewEncoder(w).Encode(&HttpResponse{
			Data:    c.Data,
			Code:    c.Code,
			Message: c.Message,
		})
	})

	http.ListenAndServe(":"+port, nil)

	select {}
}
