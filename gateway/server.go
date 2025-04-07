package gateway

import (
	"encoding/json"
	"net/http"
	"time"

	e "github.com/truc-engine/truc"
	"github.com/truc-engine/truc/util"
)

type HttpResponse struct {
	Data    *any   `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type JsonContext struct {
	Url      string
	Request  any
	Response any
	Id       string
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

		requestContext := JsonContext{
			Url:     url,
			Request: payload,
			Id:      util.NewUuidV7(),
		}
		requestJsonContext, err := json.Marshal(requestContext)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		natResponse, err := engine.Nats.Request(url, requestJsonContext, 600*time.Second)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseContext := new(e.Res[any])
		err = json.Unmarshal(natResponse.Data, &responseContext)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if responseContext.StatusCode == 0 {
			responseContext.StatusCode = 500
			responseContext.Code = "SERVER_ERROR"
			responseContext.Data = nil
			responseContext.Message = "Internal Server Error"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseContext.StatusCode)
		json.NewEncoder(w).Encode(&HttpResponse{
			Data:    responseContext.Data,
			Code:    responseContext.Code,
			Message: responseContext.Message,
		})
	})

	http.ListenAndServe(":"+port, nil)

	select {}
}
