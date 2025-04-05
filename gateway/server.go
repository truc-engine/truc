package gateway

import (
	"encoding/json"
	"net/http"
	"time"

	e "github.com/truc-engine/truc/engine"
	"github.com/truc-engine/truc/util"
)

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

		context := C{
			Url:     url,
			Request: payload,
			Id:      util.NewUuidV7(),
		}
		data, err := json.Marshal(context)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err := engine.Nats.Request(url, data, 600*time.Second)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c := new(ContextResponse)
		err = json.Unmarshal(res.Data, &c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	})

	http.ListenAndServe(":"+port, nil)
}
