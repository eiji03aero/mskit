package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ordersvc "github.com/eiji03aero/mskit/examples/ftgo/pkg/order"
)

func New(svc ordersvc.Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/order", order(svc))

	return mux
}

func order(svc ordersvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			params := ordersvc.CreateOrderParams{}
			err = json.Unmarshal(body, &params)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			id, err := svc.CreateOrder(&params)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Println("order created: ", id)
			return
		}
	}
}
