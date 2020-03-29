package restaurant

import (
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type Restaurant struct {
	mskit.BaseAggregate
	Name           string         `json:"name"`
	RestaurantMenu RestaurantMenu `json:"restaurant_menu"`
}

func (r *Restaurant) Validate() (errs []error) {
	return
}

func (r *Restaurant) Process(cmd interface{}) (mskit.Events, error) {
	switch c := cmd.(type) {
	case CreateRestaurant:
		return r.processCreateRestaurant(c)
	default:
		return mskit.Events{}, errbdr.NewErrUnknownParams(r.Process, c)
	}
}

func (r *Restaurant) processCreateRestaurant(cmd CreateRestaurant) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Restaurant{},
		RestaurantCreated{
			Id:             cmd.Id,
			Name:           cmd.Name,
			RestaurantMenu: cmd.RestaurantMenu,
		},
	)

	return events, nil
}

func (r *Restaurant) Apply(event interface{}) error {
	switch e := event.(type) {
	case RestaurantCreated:
		return r.applyRestaurantCreated(e)
	default:
		return errbdr.NewErrUnknownParams(r.Apply, e)
	}
}

func (r *Restaurant) applyRestaurantCreated(event RestaurantCreated) error {
	r.Id = event.Id
	r.Name = event.Name
	r.RestaurantMenu = event.RestaurantMenu

	return nil
}
