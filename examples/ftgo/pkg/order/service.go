package order

import (
	orderdmn "order/domain/order"
	restaurantdmn "order/domain/restaurant"

	"github.com/eiji03aero/mskit"
)

type Service interface {
	CreateOrder(params orderdmn.CreateOrder) (id string, err error)
	GetOrder(id string) (order *orderdmn.Order, err error)
	GetOrderTotal(id string) (total int, err error)
	RejectOrder(cmd orderdmn.RejectOrder) (err error)

	CreateRestaurant(restaurantdmn.Restaurant) (err error)
	GetRestaurant(id string) (restaurant *restaurantdmn.Restaurant, err error)

	InjectSagaManagers(
		createOrderSaga mskit.SagaManager,
	)
}
