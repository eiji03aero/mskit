package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"

	consumerroot "consumer"
	consumerdmn "consumer/domain/consumer"

	"github.com/eiji03aero/mskit/utils/logger"
)

func New(svc consumerroot.Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/consumers", consumers(svc))
	mux.Handle("/consumers/", consumersMember(svc))

	return mux
}

func consumers(svc consumerroot.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			logger.Println("POST /consumers")

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			cmd := consumerdmn.CreateConsumer{}
			err = json.Unmarshal(body, &cmd)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			id, err := svc.CreateConsumer(cmd)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			w.Write([]byte(id))
			return
		}
	}
}

func consumersMember(svc consumerroot.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)

		switch r.Method {
		case "GET":
			logger.Println("GET /consumers/")

			consumer, err := svc.GetConsumer(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			consumerJson, err := json.Marshal(consumer)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(consumerJson)
		}
	}
}
