package createorder

import (
	"order"

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
	svc order.Service,
	opxy order.OrderProxy,
	cpxy order.ConsumerProxy,
) mskit.SagaManager {
	c := &client{
		service:       svc,
		orderProxy:    opxy,
		consumerProxy: cpxy,
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
		Build()
	if err != nil {
		panic(err)
	}

	return mskit.NewSagaManager(
		definition,
		repository,
	)
}

type client struct {
	service       order.Service
	orderProxy    order.OrderProxy
	consumerProxy order.ConsumerProxy
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
