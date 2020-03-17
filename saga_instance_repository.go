package mskit

// SagaInstanceRepository defines interface to persist SagaInstance
type SagaInstanceRepository interface {
	Save(SagaInstance) error
	Load(string, SagaInstance) (*SagaInstance, error)
}
