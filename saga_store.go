package mskit

// SagaStore defines interface to persist SagaInstance
type SagaStore interface {
	Save(SagaInstance) error
	Load(string, *SagaInstance) error
}
