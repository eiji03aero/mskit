# ftgo
- from: https://www.manning.com/books/microservices-patterns

# Todo

# Services
## OrderService
### Aggregates
- Order
  - properties
    - orderLineItems []OrderLineItem
    - state OrderState
    - deliveryInformation DeliveryInformation
    - paymentInformation PaymentInformation

    - consumerId string
    - restaurantId string
    - orderMinimum Money
### ValueObject
- OrderState
  - enum
    - unknown, approved, approval_pending, canceled, cancel_pending, rejected, revision_pending
- OrderLineItems
  - properties
    - lineItems []OrderLineItem
- OrderLineItem
  - properties
    - menuItemId string
    - quantity int
- DeliveryInformation
  - properties
    - address Address
- PaymentInformation
  - properties
    - token string
- Address
  - properties
    - zipCode string

## KitchenService
### Aggregates
- Ticket
  - properties
    - state TicketState
    - previousState TicketState
    - restaurantId string
    - ticketLineItems TicketLineItems
    - readyBy time
    - acceptTime time
    - preparingTime time
    - pickedUpTime time
    - readyForPickupTime time

### ValueObject
- TicketState
  - enum
    - create_pending, awaiting_acceptance, accepted, preparing, ready_for_pickup, picked_up, cancel_pending, cancelled, revision_pending,
- TicketLineItems
  - properties
    - lineItems []TicketLineItem
- TicketLineItem
  - properties
    quantity int
    menuItemId string

# Common
- Money

# Apis
## create order
- OrderService#createOrder
  - arguments
    - CreateOrder{}
      - consumerId
      - restaurantId
      - lineItems []OrderMenuItem
  - flow
    - verify if restaurant exists
      - if not returns error
    - verify menuItems
      - if not returns error
    - repo.ExecuteCommand
    - create CreateOrderSaga and execute
        step()
        .withCompensation(orderService.reject, CreateOrderSagaState::makeRejectOrderCommand)
      .step()
        .invokeParticipant(consumerService.validateOrder, CreateOrderSagaState::makeValidateOrderByConsumerCommand)
      .step()
        .invokeParticipant(kitchenService.create, CreateOrderSagaState::makeCreateTicketCommand)
        .onReply(CreateTicketReply.class, CreateOrderSagaState::handleCreateTicketReply)
        .withCompensation(kitchenService.cancel, CreateOrderSagaState::makeCancelCreateTicketCommand)
      .step()
        .invokeParticipant(accountingService.authorize, CreateOrderSagaState::makeAuthorizeCommand)
      .step()
        .invokeParticipant(kitchenService.confirmCreate, CreateOrderSagaState::makeConfirmCreateTicketCommand)
      .step()
        .invokeParticipant(orderService.approve, CreateOrderSagaState::makeApproveOrderCommand)
      .build();
