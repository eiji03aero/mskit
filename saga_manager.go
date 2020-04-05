package mskit

import (
	"encoding/json"
	"reflect"

	"github.com/eiji03aero/mskit/utils/errbdr"
	"github.com/eiji03aero/mskit/utils/logger"
)

type sagaManager struct {
	resultChannel   chan *SagaStepResult
	repository      *SagaRepository
	sagaDefinition  *SagaDefinition
	sagaStateStruct interface{}
}

// SagaManager is an interface that manages saga operation of certain type
type SagaManager interface {
	Create(sagaState interface{}) (si *SagaInstance, err error)
	Subscribe()
}

// NewSagaManager creates SagaManager object
func NewSagaManager(sr *SagaRepository, sd *SagaDefinition, sss interface{}) *sagaManager {
	return &sagaManager{
		resultChannel:   make(chan *SagaStepResult),
		repository:      sr,
		sagaDefinition:  sd,
		sagaStateStruct: sss,
	}
}

// Create creates new SagaInstance and triggers flow
func (sm *sagaManager) Create(sagaState interface{}) (si *SagaInstance, err error) {
	logger.Println(
		logger.HiBlueString("Create Saga"),
		si,
	)

	si, err = NewSagaInstance()
	if err != nil {
		return
	}

	si.Data = sagaState
	err = sm.repository.Save(si)
	if err != nil {
		return
	}

	err = sm.executeStep(si)
	if err != nil {
		return
	}

	err = sm.repository.Update(si)
	if err != nil {
		return
	}

	return
}

// Subscribe starts subscribing to step results
func (sm *sagaManager) Subscribe() {
	for result := range sm.resultChannel {
		sm.processResult(result)
	}
}

func (sm *sagaManager) processResult(result *SagaStepResult) (err error) {
	si := &SagaInstance{}

	err = sm.repository.Load(result.Id, si)
	if err != nil {
		return
	}

	err = sm.restoreData(si)
	if err != nil {
		return
	}

	err = si.processResult(result)
	if err != nil {
		return
	}

	err = sm.executeStep(si)
	if err != nil {
		return
	}

	err = sm.repository.Update(si)
	if err != nil {
		return
	}

	return
}

func (sm *sagaManager) restoreData(si *SagaInstance) (err error) {
	dataStr, ok := si.Data.(string)
	if !ok {
		return errbdr.NewErrUnknownParams(sm.restoreData, si)
	}

	stateStruct := reflect.New(reflect.TypeOf(sm.sagaStateStruct)).Interface()
	err = json.Unmarshal([]byte(dataStr), stateStruct)
	if err != nil {
		return
	}

	si.Data = stateStruct
	return
}

func (sm *sagaManager) executeStep(si *SagaInstance) (err error) {
	for {
		if si.checkFinishState(sm.sagaDefinition.Len()) {
			break
		}

		step := sm.sagaDefinition.Get(si.StepIndex)

		if !si.checkStepHasHandler(step) {
			logger.Println(
				logger.YellowString("SagaStep no handler to call, skipping"),
				si,
			)
			// skip if step does not have handler
			si.shiftIndex()
			continue
		}

		result := &SagaStepResult{}
		result.Id = si.Id
		result.Error = si.executeStepHandler(step)

		go sm.sendResult(result)
		break
	}

	return
}

func (sm *sagaManager) sendResult(result *SagaStepResult) {
	sm.resultChannel <- result
}
