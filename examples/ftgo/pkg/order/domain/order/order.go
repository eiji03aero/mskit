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
	TicketId            string              `json:"ticket_id"`
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
	case BeginReviseOrder:
		return o.processBeginReviseOrder(c)
	case UndoBeginReviseOrder:
		return o.processUndoBeginReviseOrder(c)
	case ConfirmReviseOrder:
		return o.processConfirmReviseOrder(c)
	case HandleTicketCreated:
		return o.processHandleTicketCreated(c)
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

func (o *Order) processBeginReviseOrder(cmd BeginReviseOrder) (events mskit.Events, err error) {
	if o.OrderState != OrderState_Approved {
		err = errbdr.NewErrUnsupportedStateTransition(o, o.OrderState)
		return
	}

	events = mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderRevisionBegan{
			Id: cmd.Id,
		},
	)

	return
}

func (o *Order) processUndoBeginReviseOrder(cmd UndoBeginReviseOrder) (events mskit.Events, err error) {
	if o.OrderState != OrderState_RevisionPending {
		err = errbdr.NewErrUnsupportedStateTransition(o, o.OrderState)
		return
	}

	events = mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		UndoOrderRevisionBegan{
			Id: cmd.Id,
		},
	)

	return
}

func (o *Order) processConfirmReviseOrder(cmd ConfirmReviseOrder) (events mskit.Events, err error) {
	if o.OrderState != OrderState_RevisionPending {
		err = errbdr.NewErrUnsupportedStateTransition(o, o.OrderState)
		return
	}

	events = mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderRevisionConfirmed{
			Id:             cmd.Id,
			OrderLineItems: cmd.OrderLineItems,
		},
	)

	return
}

func (o *Order) processHandleTicketCreated(cmd HandleTicketCreated) (events mskit.Events, err error) {
	events = mskit.NewEventsSingle(
		cmd.Id,
		Order{},
		OrderTicketIdSet{
			TicketId: cmd.TicketId,
		},
	)

	return
}

func (o *Order) Apply(event interface{}) error {
	switch e := event.(type) {
	case OrderCreated:
		return o.applyOrderCreated(e)
	case OrderRejected:
		return o.applyOrderRejected(e)
	case OrderApproved:
		return o.applyOrderApproved(e)
	case OrderRevisionBegan:
		return o.applyOrderRevisionBegan(e)
	case UndoOrderRevisionBegan:
		return o.applyUndoOrderRevisionBegan(e)
	case OrderRevisionConfirmed:
		return o.applyOrderRevisionConfirmed(e)
	case OrderTicketIdSet:
		return o.applyOrderTicketIdSet(e)
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

func (o *Order) applyOrderRevisionBegan(event OrderRevisionBegan) (err error) {
	o.OrderState = OrderState_RevisionPending
	return nil
}

func (o *Order) applyUndoOrderRevisionBegan(event UndoOrderRevisionBegan) (err error) {
	o.OrderState = OrderState_Approved
	return nil
}

func (o *Order) applyOrderRevisionConfirmed(event OrderRevisionConfirmed) (err error) {
	o.OrderState = OrderState_Approved
	o.OrderLineItems = event.OrderLineItems
	return nil
}

func (o *Order) applyOrderTicketIdSet(event OrderTicketIdSet) (err error) {
	o.TicketId = event.TicketId
	return nil
}
