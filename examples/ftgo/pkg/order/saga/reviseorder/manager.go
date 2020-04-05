package reviseorder

import (
	"order"

	"github.com/eiji03aero/mskit"
)

type client struct {
	sagaRepository  *mskit.SagaRepository
	service         order.Service
	orderProxy      order.OrderProxy
	kitchenProxy    order.KitchenProxy
	accountingProxy order.AccountingProxy
}

func NewManager(
	sagaRepository *mskit.SagaRepository,
	svc order.Service,
	opxy order.OrderProxy,
	kpxy order.KitchenProxy,
	apxy order.AccountingProxy,
) mskit.SagaManager {
	c := &client{
		sagaRepository:  sagaRepository,
		service:         svc,
		orderProxy:      opxy,
		kitchenProxy:    kpxy,
		accountingProxy: apxy,
	}

	definition, err := mskit.NewSagaDefinitionBuilder().
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.beginReviseOrderE,
			},
			mskit.SagaStepCompensationOption{
				Handler: c.beginReviseOrderC,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.beginReviseTicketE,
			},
			mskit.SagaStepCompensationOption{
				Handler: c.beginReviseTicketC,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.reviseAuthorizationE,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.confirmReviseTicketE,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.confirmReviseOrderE,
			},
		).
		Build()
	if err != nil {
		panic(err)
	}

	return mskit.NewSagaManager(
		sagaRepository,
		definition,
		state{},
	)
}

func (c *client) beginReviseOrderE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.orderProxy.BeginReviseOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	return
}

func (c *client) beginReviseOrderC(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.orderProxy.UndoBeginReviseOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	return
}

func (c *client) beginReviseTicketE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	err = c.kitchenProxy.BeginReviseTicket(order.TicketId, sagaState.OrderLineItems)
	if err != nil {
		return
	}

	return
}

func (c *client) beginReviseTicketC(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	err = c.kitchenProxy.UndoBeginReviseTicket(order.TicketId)
	if err != nil {
		return
	}

	return
}

func (c *client) reviseAuthorizationE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	err = c.accountingProxy.Authorize(order.ConsumerId)
	if err != nil {
		return
	}

	return
}

func (c *client) confirmReviseTicketE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	err = c.kitchenProxy.ConfirmReviseTicket(order.TicketId, sagaState.OrderLineItems)
	if err != nil {
		return
	}

	return
}

func (c *client) confirmReviseOrderE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.orderProxy.ConfirmReviseOrder(sagaState.OrderId, sagaState.OrderLineItems)
	if err != nil {
		return
	}

	return
}
