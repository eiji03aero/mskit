package mskit

import "log"

type sagaManager struct {
	resultChannel  chan *SagaStepResult
	repository     *SagaRepository
	sagaDefinition *SagaDefinition
}

// SagaManager is an interface that manages saga operation of certain type
type SagaManager interface {
	Create(sagaState interface{}) (si *SagaInstance, err error)
	Subscribe()
}

// NewSagaManager creates SagaManager object
func NewSagaManager(sd *SagaDefinition, sr *SagaRepository) *sagaManager {
	return &sagaManager{
		resultChannel:  make(chan *SagaStepResult),
		repository:     sr,
		sagaDefinition: sd,
	}
}

// Create creates new SagaInstance and triggers flow
func (sm *sagaManager) Create(sagaState interface{}) (si *SagaInstance, err error) {
	si, err = NewSagaInstance()
	if err != nil {
		return
	}

	log.Println("SagaManager#Create: id=", si.Id)

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

	log.Println("SagaManager#processResult: id=", si.Id)

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

func (sm *sagaManager) executeStep(si *SagaInstance) (err error) {
	for {
		if si.checkFinishState(sm.sagaDefinition.Len()) {
			log.Println("SagaManager#executeStep: finished", si.StepIndex)
			break
		}

		step := sm.sagaDefinition.Get(si.StepIndex)
		log.Println("SagaManager#executeStep: ", si, step)

		if !si.checkStepHasHandler(step) {
			log.Println("SagaManager#executeStep: skipping")
			// skip if step does not have handler
			si.shiftIndex()
			continue
		}

		log.Println("SagaManager#executeStep: gonna execute ", si)

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
