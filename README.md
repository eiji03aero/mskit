# mskit
- toolkit for microservices in go

# Todo
- create consumer service
  - validate order by consumer
- create accounting service
  - authorize order
- saga
- review all the namings
- setup tests
  - unit tests
  - add comments on exporteds
  - github action

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
  - actions []SagaAction
### SagaDefinitionBuilder
- properties
  - sagaDefinition SagaDefinition
- methods
  - Step()
  - WithCompensation()
  - InvokeParticipant()
  - OnReply()
  - Build() (SagaDefinition, error)
### SagaAction
- properties
  - OnCompensation func()
  - OnReply func()
  - OnInvokeParticipant func()
- methods
  - validate() error
### SagaInstance
- properties
  - SagaState enum [Processing, Aborted, Done]
  - Data interface{}
### SagaManager
- methods
  - Create(SagaDefinition) (SagaInstance)
  - Subscribe

## interfaces
### Saga
- methods
  - GetDefinition () (SagaDefinition)

## concerns
- instance that runs across deployment could cause error if saga definition had update
