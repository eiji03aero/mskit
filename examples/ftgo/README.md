# ftgo
- from: https://www.manning.com/books/microservices-patterns

# Services
- OrderService
- RestaurantService

# Apis
## create order
- OrderService#createOrder
  - order := Order.New
  - events, err := order.process(cmd CreateOrder)
  - err := order.apply(events)
  - repository.save(events)

# Components
## BaseAggregate
- properties
  - ID string
- methods
  - Process(cmd Command) (events, err)
  - Apply(event) (err)
  - Type() string

## Event
- properties
  - ID string
  - Type string
  - AggregateID string
  - AggregateType string
  - Data interface{}

## Repository
- methods
  - Save(events []Event) err
  - Load(id string) (BaseAggregate, err)
