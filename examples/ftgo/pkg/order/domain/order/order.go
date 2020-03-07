package order

import (
	"errors"
	"fmt"

	"github.com/eiji03aero/mskit"
)

type Order struct {
	mskit.BaseAggregate
	OrderState          OrderState          `json:"order_state"`
	PaymentInformation  PaymentInformation  `json:"payment_information"`
	DeliveryInformation DeliveryInformation `json:"delivery_information"`
	OrderLineItems      OrderLineItems      `json:"order_line_items"`
}

func (o *Order) Validate() (errs []error) {
	if o.OrderLineItems.Len() < 1 {
		errs = append(errs, errors.New("quantity of order line items not enough"))
	}

	return errs
}

func (o *Order) Process(cmd interface{}) (mskit.Events, error) {
	switch c := cmd.(type) {
	case CreateOrder:
		return o.processCreateOrder(c)
	default:
		return nil, errors.New("not imp in Process")
	}
}

func (o *Order) Apply(event interface{}) error {
	switch e := event.(type) {
	case OrderCreated:
		return o.applyOrderCreated(e)
	default:
		return errors.New(fmt.Sprintf("not implemented in Apply: %v", e))
	}
}

func (o *Order) processCreateOrder(cmd CreateOrder) (mskit.Events, error) {
	// validation has to be here, return error if bad

	events := mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderCreated{
			Id:                  cmd.Id,
			PaymentInformation:  cmd.PaymentInformation,
			DeliveryInformation: cmd.DeliveryInformation,
			OrderLineItems:      cmd.OrderLineItems,
		},
	)

	return events, nil
}

func (o *Order) applyOrderCreated(event OrderCreated) (err error) {
	o.OrderState = OrderState_ApprovalPending
	o.Id = event.Id

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
