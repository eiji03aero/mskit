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
	ApproveOrder(cmd orderdmn.ApproveOrder) (err error)
	ReviseOrder(cmd orderdmn.ReviseOrder) (err error)
	BeginReviseOrder(cmd orderdmn.BeginReviseOrder) (err error)
	UndoBeginReviseOrder(cmd orderdmn.UndoBeginReviseOrder) (err error)
	ConfirmReviseOrder(cmd orderdmn.ConfirmReviseOrder) (err error)

	HandleTicketCreated(cmd orderdmn.HandleTicketCreated) (err error)

	CreateRestaurant(restaurantdmn.Restaurant) (err error)
	GetRestaurant(id string) (restaurant *restaurantdmn.Restaurant, err error)

	InjectSagaManagers(
		createOrderSagaManager mskit.SagaManager,
		reviseOrderSagaManager mskit.SagaManager,
	)
}
