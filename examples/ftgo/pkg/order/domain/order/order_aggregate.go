package order

import (
	"fmt"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type Order struct {
	mskit.BaseAggregate
	ConsumerId          string              `json:"consumer_id"`
	RestaurantId        string              `json:"restaurant_id"`
	OrderState          OrderState          `json:"order_state"`
	PaymentInformation  PaymentInformation  `json:"payment_information"`
	DeliveryInformation DeliveryInformation `json:"delivery_information"`
	OrderLineItems      OrderLineItems      `json:"order_line_items"`
}

func (o *Order) Validate() (errs []error) {
	if o.OrderLineItems.Len() < 1 {
		errs = append(errs, fmt.Errorf("quantity of order line items not enough"))
	}

	return errs
}

func (o *Order) Process(cmd interface{}) (mskit.Events, error) {
	switch c := cmd.(type) {
	case CreateOrder:
		return o.processCreateOrder(c)
	case RejectOrder:
		return o.processRejectOrder(c)
	case ApproveOrder:
		return o.processApproveOrder(c)
	default:
		return nil, errbdr.NewErrUnknownParams(o.Process, c)
	}
}

func (o *Order) processCreateOrder(cmd CreateOrder) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderCreated{
			Id:                  cmd.Id,
			ConsumerId:          cmd.ConsumerId,
			RestaurantId:        cmd.RestaurantId,
			PaymentInformation:  cmd.PaymentInformation,
			DeliveryInformation: cmd.DeliveryInformation,
			OrderLineItems:      cmd.OrderLineItems,
		},
	)

	return events, nil
}

func (o *Order) processRejectOrder(cmd RejectOrder) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderRejected{
			Id: cmd.Id,
		},
	)

	return events, nil
}

func (o *Order) processApproveOrder(cmd ApproveOrder) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderApproved{
			Id: cmd.Id,
		},
	)

	return events, nil
}

func (o *Order) Apply(event interface{}) error {
	switch e := event.(type) {
	case OrderCreated:
		return o.applyOrderCreated(e)
	case OrderRejected:
		return o.applyOrderRejected(e)
	case OrderApproved:
		return o.applyOrderApproved(e)
	default:
		return errbdr.NewErrUnknownParams(o.Apply, e)
	}
}

func (o *Order) applyOrderCreated(event OrderCreated) (err error) {
	o.OrderState = OrderState_ApprovalPending
	o.Id = event.Id
	o.RestaurantId = event.RestaurantId
	o.ConsumerId = event.ConsumerId

	err = o.PaymentInformation.Merge(event.PaymentInformation)
	if err != nil {
		return err
	}

	err = o.DeliveryInformation.Merge(event.DeliveryInformation)
	if err != nil {
		return err
	}

	err = o.OrderLineItems.Merge(event.OrderLineItems)
	if err != nil {
		return err
	}

	return nil
}

func (o *Order) applyOrderRejected(event OrderRejected) (err error) {
	o.OrderState = OrderState_Rejected
	return nil
}

func (o *Order) applyOrderApproved(event OrderApproved) (err error) {
	o.OrderState = OrderState_Approved
	return nil
}
