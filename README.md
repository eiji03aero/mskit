# mskit
- toolkit for microservices in go

# Todo
- review namings
  - none for now

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
### Saga
- properties
  - Definition \*SagaDefinition
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
  - InvokeParticipant()
  - OnReply()
  - WithCompensation()
  - Build() (SagaDefinition, error)
### SagaStepInvokeParticipantOption
- handler func()
### SagaStepReplyOption
- handler func()
### SagaStepCompensationOption
- handler func()
### SagaStep
- properties
  - invokeParticipantHandler func()
  - replyHandler func()
  - compensationHandler func()
- methods
  - Validate() error
  - Configure(opts ...interface{})
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
