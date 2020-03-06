# mskit
- toolkit for microservices in go

# Todo
- probably publish should be handled by repository
- add domain event pug/sub on restaurant, order
  - Restaurant created event
- add default options on pub/sub/rpcs
  - deal with how to configure as well. should be with dedicated method to set options
- try to deal with basic process/apply usecase, create helper
- setup tests
  - unit tests
  - github action
- should try to make the pub/sub and rpc apis easier
- update aggregate
  - version
  - snapshot
- review all the namings

- should ponder how tables are initialized
  - probably it should just provide sql

# Features
- Event sourcing
- Pub/Sub
  - EventPublisher
  - EventConsumer
- RPC with other service
- Saga

# Components
## BaseAggregate
- properties
  - ID string
- methods
  - Process(cmd Command) (events, err)
  - Apply(event) (err)

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

## EventBus(rabbitmq)
- properties
  - conn rabbitmq.Connection
- methods
  - NewPublisher
  - NewConsumer
  - NewRPCServer
  - NewRPCClient
