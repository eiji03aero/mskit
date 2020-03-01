# mskit
- toolkit for microservices in go

# Todo
- setup tests
  - unit tests
  - github action
- should try to make the pub/sub, rpc apis easier ...
- should ponder how tables are initialized
  - probably it should just provide sql
- update aggregate
  - version
  - snapshot

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
