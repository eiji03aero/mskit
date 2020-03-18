package mskit

// SagaStore defines interface to persist SagaInstance
type SagaRepository struct {
	sagaStore SagaStore
}

// NewRepository creates and returns new SagaRepository
func NewSagaRepository(ss SagaStore) *SagaRepository {
	return &SagaRepository{
		sagaStore: ss,
	}
}

// Save persists SagaInstance
func (sr *SagaRepository) Save(si SagaInstance) error {
	return sr.sagaStore.Save(si)
}

// Load loads up data into SagaInstance
func (sr *SagaRepository) Load(id string, si *SagaInstance) error {
	return sr.sagaStore.Load(id, si)
}
