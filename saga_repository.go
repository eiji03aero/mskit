package mskit

// SagaStore defines interface to persist SagaInstance
type SagaRepository struct {
	sagaStore SagaStore
}

// NewSagaRepository creates and returns new SagaRepository
func NewSagaRepository(ss SagaStore) *SagaRepository {
	return &SagaRepository{
		sagaStore: ss,
	}
}

// Save persists SagaInstance
func (sr *SagaRepository) Save(si *SagaInstance) error {
	return sr.sagaStore.Save(si)
}

// Update saves updates of SagaInstance
func (sr *SagaRepository) Update(si *SagaInstance) error {
	return sr.sagaStore.Update(si)
}

// Load loads up data into SagaInstance
func (sr *SagaRepository) Load(id string, si *SagaInstance) error {
	return sr.sagaStore.Load(id, si)
}
