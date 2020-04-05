package mskit

import (
	"github.com/eiji03aero/mskit/utils/logger"
)

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
func (sr *SagaRepository) Save(si *SagaInstance) (err error) {
	logger.Println(
		logger.HiBlueString("Save SagaInstance"),
		si,
	)

	err = sr.sagaStore.Save(si)
	if err != nil {
		return
	}
	return
}

// Update saves updates of SagaInstance
func (sr *SagaRepository) Update(si *SagaInstance) (err error) {
	logger.Println(
		logger.HiBlueString("Update SagaInstance"),
		si,
	)

	err = sr.sagaStore.Update(si)
	if err != nil {
		return
	}

	return
}

// Load loads up data into SagaInstance
func (sr *SagaRepository) Load(id string, si *SagaInstance) (err error) {
	logger.HiBlueString(
		logger.HiBlueString("Load SagaInstance"),
		si,
	)

	err = sr.sagaStore.Load(id, si)
	if err != nil {
		return
	}

	return
}
