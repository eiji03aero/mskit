# mskit
- toolkit for microservices in go

# Todo
- review namings
  - none for now

- cli
  - create service
    - skelton structure
      - cmd
      - domain
      - service
        - service.go
  - generate aggregate
  - generate saga
  - generate interface file on root

- refactor
  - cli/template related

- improve logging
  - add level
  - add color

- saga
  - createOrder saga
    - add the initial rejectOrder step
      - orderservice proxy
    - validate order by consumer
      - find consumer by id
        - if not found, return error
      - call consumer.validateOrder
    - create accounting service
      - authorize order

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

# Saga
## flow
### preparation
- create instance of SagaDefinition
  - on ram
- create instance of SagaManager (for each saga)
  - on ram
  - create consumer to listen for command response
### when create
- invoke manager to create instance of SagaInstance and trigger
  - create(SagaDefinition, [other required attributes])
  - run the first step
  - persists SagaInstance
### when invoked by consumer
- wait for response and invoke saga
  - listen for saga
- if response was fail, invoke compensations

## components
### SagaDefinition
- properties
  - steps []SagaStep
- methods
  - addStep()
### SagaDefinitionBuilder
- properties
  - sagaDefinition SagaDefinition
- methods
  - Step(opts ...interface{}) (\*SagaDefinitionBuilder)
  - Build() (SagaDefinition, error)

### SagaStepExecuteOption
- Handler func(sagaInstance interface{}) (error)
### SagaStepCompensationOption
- Handler func(sagaInstance interface{}) (error)
### SagaStep
- properties
  - executeHandler func()
  - compensationHandler func()
- methods
  - Validate() error
  - Configure(opts ...interface{})
### SagaStepResult
- properties
  - Id string
  - Error error

### SagaInstance
- properties
  - Id string
  - SagaState enum [Processing, Aborted, Done]
  - Data interface{}

### SagaManager
- properties
  - repository SagaRepository
  - saga Saga
- methods
  - Create(SagaDefinition) (SagaInstance)
  - Subscribe
