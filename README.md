# mskit
- toolkit for microservices in go

# Todo
- create consumer service
  - validate order by consumer
- create accounting service
  - authorize order
- saga
- review all the namings
  - NewRPCEndpoint instead of NewRPCServer?
- setup tests
  - unit tests
  - add comments on exporteds
  - github action
- move docker related files into sub directory

- ponder on how repositories should be managed
  - unify sagainstancerepo and event repo into one?
  - keep them both separate?
  - btw naming on sagastore is not consistent, while eventrepo has both eventstore and repository
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
### SagaStepReplyOption
### SagaStepCompensationOption
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
  - repository SagaInstanceRepository
  - saga Saga
- methods
  - Create(SagaDefinition) (SagaInstance)
  - Subscribe
