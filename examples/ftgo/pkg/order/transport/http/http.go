package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"

	orderdmn "order/domain/order"
	ordersvc "order/service"
)

func New(svc ordersvc.Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/orders", orders(svc))
	mux.Handle("/orders/", ordersMember(svc))

	mux.Handle("/restaurants/", restaurantsMember(svc))

	return mux
}

func orders(svc ordersvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
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

func ordersMember(svc ordersvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)

		switch r.Method {
		case "GET":
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
		}
	}
}

func restaurantsMember(svc ordersvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)

		switch r.Method {
		case "GET":
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
