# mskit
- toolkit for microservices in go

# Todo
- add kitchen service
  - create mongo eventstore
  - no protobuffer

## TBD
- should ponder how tables are initialized
  - probably it should just provide sql
- version on aggregate

# Features
- Event sourcing
- Saga
- Pub/Sub
  - EventPublisher
- RPC with other service
  - RPCClient

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
