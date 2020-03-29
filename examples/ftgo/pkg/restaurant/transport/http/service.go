package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	restaurantdmn "restaurant/domain/restaurant"
	restaurantsvc "restaurant/service"

	"github.com/eiji03aero/mskit/utils/logger"
)

func New(svc restaurantsvc.Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/restaurants", restaurants(svc))

	return mux
}

func restaurants(svc restaurantsvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			logger.Println("POST /restaurants")

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			cmd := restaurantdmn.CreateRestaurant{}
			err = json.Unmarshal(body, &cmd)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			id, err := svc.CreateRestaurant(cmd)
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
