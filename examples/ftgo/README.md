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

## RestaurantService

# Common
- Money

# Apis
## create order
- OrderService#createOrder
  - order := Order.New
  - events, err := order.process(cmd CreateOrder)
  - err := order.apply(events)
  - repository.save(events)
