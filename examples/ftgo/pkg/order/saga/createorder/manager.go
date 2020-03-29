package createorder

import (
	orderroot "order"

	"github.com/eiji03aero/mskit"
)

// create CreateOrderSaga and execute
//     .step()
//       .invokeParticipant(kitchenService.create, CreateOrderSagaState::makeCreateTicketCommand)
//       .onReply(CreateTicketReply.class, CreateOrderSagaState::handleCreateTicketReply)
//       .withCompensation(kitchenService.cancel, CreateOrderSagaState::makeCancelCreateTicketCommand)
//     .step()
//       .invokeParticipant(accountingService.authorize, CreateOrderSagaState::makeAuthorizeCommand)
//     .step()
//       .invokeParticipant(kitchenService.confirmCreate, CreateOrderSagaState::makeConfirmCreateTicketCommand)
//     .step()
//       .invokeParticipant(orderService.approve, CreateOrderSagaState::makeApproveOrderCommand)
//     .build();
func NewManager(
	repository *mskit.SagaRepository,
	svc orderroot.Service,
	opxy orderroot.OrderProxy,
	cpxy orderroot.ConsumerProxy,
	kpxy orderroot.KitchenProxy,
) mskit.SagaManager {
	c := &client{
		service:       svc,
		orderProxy:    opxy,
		consumerProxy: cpxy,
		kitchenProxy:  kpxy,
	}

	definition, err := mskit.NewSagaDefinitionBuilder().
		Step(
			mskit.SagaStepCompensationOption{
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
		).
		Build()
	if err != nil {
		panic(err)
	}

	return mskit.NewSagaManager(
		repository,
		definition,
		state{},
	)
}

type client struct {
	service       orderroot.Service
	orderProxy    orderroot.OrderProxy
	consumerProxy orderroot.ConsumerProxy
	kitchenProxy  orderroot.KitchenProxy
}

func (c *client) rejectOrderC(ss interface{}) (err error) {
	sagaState, err := assertStruct(ss)
	if err != nil {
		return
	}

	err = c.orderProxy.RejectOrder(sagaState.OrderId)

	return
}

func (c *client) validateOrderE(ss interface{}) (err error) {
	sagaState, err := assertStruct(ss)
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

func (c *client) createTicketE(ss interface{}) (err error) {
	sagaState, err := assertStruct(ss)
	if err != nil {
		return
	}

	order, err := c.service.GetOrder(sagaState.OrderId)
	if err != nil {
		return
	}

	_, err = c.kitchenProxy.CreateTicket(
		order.RestaurantId,
		order.OrderLineItems.LineItems,
	)
	if err != nil {
		return
	}

	// TODO: update ticketId on somewhere. on order?

	return
}
