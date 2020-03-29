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
	err = sr.sagaStore.Save(si)
	if err != nil {
		return
	}

	logger.PrintFuncCall(sr.Save, si)
	return
}

// Update saves updates of SagaInstance
func (sr *SagaRepository) Update(si *SagaInstance) (err error) {
	err = sr.sagaStore.Update(si)
	if err != nil {
		return
	}

	logger.PrintFuncCall(sr.Save, si)
	return
}

// Load loads up data into SagaInstance
func (sr *SagaRepository) Load(id string, si *SagaInstance) (err error) {
	err = sr.sagaStore.Load(id, si)
	if err != nil {
		return
	}

	logger.PrintFuncCall(sr.Load, si)
	return
}
