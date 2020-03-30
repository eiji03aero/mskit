# mskit
- toolkit for microservices in go

# Todo
- review namings
  - none for now

- cli
  - generate saga

- improve logging
  - add level

- saga
  - createOrder saga

- refactor
  - try simplify the initialization in cmd/app :(
  - logging
    - should add logging on
      - publish, consume, rpcendpoint, rpcclient
  - should think about how eventstore and sagastore dealing with restoring data from json

- rpc
  - add timeout

- refactor services
  - align the structure

- setup tests
  - unit tests
  - add comments on exporteds
  - github action

- update aggregate
  - version
  - snapshot

- udpate go to 1.14

# Concerns
- better implementation for saga?
  - should remain to be flexible (capable of using messaging/http/grpc)

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
  - Id string
- methods
  - Process(cmd Command) (events, err)
  - Apply(event) (err)

## Event
- properties
  - Id string
  - Type string
  - AggregateId string
  - AggregateType string
  - Data interface{}

## EventRepository
- methods
  - Save(events []Event) err
  - Load(id string) (BaseAggregate, err)

## EventBus(rabbitmq)
- properties
  - conn rabbitmq.Connection
- methods
  - NewPublisher
  - NewConsumer
  - NewRPCEndpoint
  - NewRPCClient

## SagaDefinition
- properties
  - steps []SagaStep
- methods
  - addStep()
## SagaDefinitionBuilder
- properties
  - sagaDefinition SagaDefinition
- methods
  - Step(opts ...interface{}) (\*SagaDefinitionBuilder)
  - Build() (SagaDefinition, error)

## SagaStepExecuteOption
- Handler func(sagaInstance interface{}) (error)
## SagaStepCompensationOption
- Handler func(sagaInstance interface{}) (error)
## SagaStep
- properties
  - executeHandler func()
  - compensationHandler func()
- methods
  - Validate() error
  - Configure(opts ...interface{})
## SagaStepResult
- properties
  - Id string
  - Error error

## SagaInstance
- properties
  - Id string
  - SagaState enum [Processing, Aborted, Done]
  - Data interface{}

## SagaManager
- properties
  - repository SagaRepository
  - saga Saga
- methods
  - Create(SagaDefinition) (SagaInstance)
  - Subscribe

# CLI

```sh
# initialize service
$ mskit init order

# cd into service directory
$ cd order

# generate aggregate
$ mskit generate aggregate Order

# generate proxy
$ mskit generate proxy Kitchen

# generate rpcendpoint
$ mskit generate rpcendpoint

# generate publisher
$ mskit generate publisher

# generate consumer
$ mskit generate consumer
```
