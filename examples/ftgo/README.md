# ftgo
- from: https://www.manning.com/books/microservices-patterns

# Todo
- update Order to include all properties and logics related
- add common events
  - not implemented events/commands

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
  - order := Order.New
  - events, err := order.process(cmd CreateOrder)
  - err := order.apply(events)
  - repository.save(events)
