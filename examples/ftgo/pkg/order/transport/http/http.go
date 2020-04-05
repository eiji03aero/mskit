package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"

	"order"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/utils/logger"
)

func New(svc order.Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/orders", orders(svc))
	mux.Handle("/orders/", ordersMember(svc))

	mux.Handle("/restaurants/", restaurantsMember(svc))

	return mux
}

func orders(svc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			logger.Println("POST /orders")

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			params := orderdmn.CreateOrder{}
			err = json.Unmarshal(body, &params)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			id, err := svc.CreateOrder(params)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			w.Write([]byte(id))
		}
	}
}

func ordersMember(svc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)

		switch r.Method {
		case "GET":
			logger.Println("GET /orders/", id)

			order, err := svc.GetOrder(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			orderJson, err := json.Marshal(order)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(orderJson)

		case "PATCH":
			logger.Println("PATCH /orders/", id)

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			cmd := orderdmn.ReviseOrder{Id: id}
			err = json.Unmarshal(body, &cmd)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			err = svc.ReviseOrder(cmd)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}

func restaurantsMember(svc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)

		switch r.Method {
		case "GET":
			logger.Println("GET /restaurants/")

			restaurant, err := svc.GetRestaurant(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			restaurantJson, err := json.Marshal(restaurant)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(restaurantJson)
		}
	}
}
