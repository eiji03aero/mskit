package createorder

import (
	orderroot "order"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit"
)

type client struct {
	sagaRepository  *mskit.SagaRepository
	service         orderroot.Service
	orderProxy      orderroot.OrderProxy
	consumerProxy   orderroot.ConsumerProxy
	kitchenProxy    orderroot.KitchenProxy
	accountingProxy orderroot.AccountingProxy
}

func NewManager(
	sagaRepository *mskit.SagaRepository,
	svc orderroot.Service,
	opxy orderroot.OrderProxy,
	cpxy orderroot.ConsumerProxy,
	kpxy orderroot.KitchenProxy,
	apxy orderroot.AccountingProxy,
) mskit.SagaManager {
	c := &client{
		sagaRepository:  sagaRepository,
		service:         svc,
		orderProxy:      opxy,
		consumerProxy:   cpxy,
		kitchenProxy:    kpxy,
		accountingProxy: apxy,
	}

	definition, err := mskit.NewSagaDefinitionBuilder().
		Step(
			mskit.SagaStepCompensateOption{
				Handler: c.rejectOrderC,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.validateOrderE,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.createTicketE,
			},
			mskit.SagaStepCompensateOption{
				Handler: c.createTicketC,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.authorizeConsumerE,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.confirmCreateTicketE,
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.approveOrderE,
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

func (c *client) rejectOrderC(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.orderProxy.RejectOrder(sagaState.OrderId)

	return
}

func (c *client) validateOrderE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	total, err := c.service.GetOrderTotal(sagaState.OrderId)
	if err != nil {
		return
	}

	err = c.consumerProxy.ValidateOrder(order.ConsumerId, total)
	if err != nil {
		return
	}

	return
}

func (c *client) createTicketE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	ticketId, err := c.kitchenProxy.CreateTicket(
		order.RestaurantId,
		order.OrderLineItems,
	)
	if err != nil {
		return
	}

	sagaState.TicketId = ticketId
	si.Data = sagaState
	err = c.sagaRepository.Update(si)
	if err != nil {
		return
	}

	err = c.service.HandleTicketCreated(orderdmn.HandleTicketCreated{
		Id:       sagaState.OrderId,
		TicketId: sagaState.TicketId,
	})
	if err != nil {
		return
	}

	return
}

func (c *client) createTicketC(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.kitchenProxy.CancelTicket(sagaState.TicketId)
	if err != nil {
		return
	}

	return
}

func (c *client) authorizeConsumerE(si *mskit.SagaInstance) (err error) {
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

func (c *client) confirmCreateTicketE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.kitchenProxy.ConfirmTicket(sagaState.TicketId)
	if err != nil {
		return
	}

	return
}

func (c *client) approveOrderE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	err = c.orderProxy.ApproveOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	return
}
