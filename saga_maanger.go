package mskit

// SagaManager is a struct that manages saga operation
type SagaManager struct {
	repository interface{}
	saga       Saga
}

// NewSagaManager creates SagaManager object
func NewSagaManager(saga Saga, repository interface{}) *SagaManager {
	return &SagaManager{
		repository: repository,
		saga:       saga,
	}
}

// Create creates new SagaInstance and triggers flow
func (sm *SagaManager) Create(sagaDefinition SagaDefinition) *SagaInstance {
	return nil
}

// Subscribe starts subscribing to channels
func (sm *SagaManager) Subscribe() error {
	return nil
}
