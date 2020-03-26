package mskit

// SagaManager is a struct that manages saga operation of certain type
type SagaManager struct {
	resultChannel  chan *SagaStepResult
	repository     *SagaRepository
	sagaDefinition *SagaDefinition
}

// NewSagaManager creates SagaManager object
func NewSagaManager(sd *SagaDefinition, sr *SagaRepository) *SagaManager {
	return &SagaManager{
		resultChannel:  make(chan *SagaStepResult),
		repository:     sr,
		sagaDefinition: sd,
	}
}

// Create creates new SagaInstance and triggers flow
func (sm *SagaManager) Create() (si *SagaInstance, err error) {
	si, err = NewSagaInstance()
	if err != nil {
		return
	}

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
func (sm *SagaManager) Subscribe() {
	go func() {
		for result := range sm.resultChannel {
			sm.processResult(result)
		}
	}()
}

func (sm *SagaManager) processResult(result *SagaStepResult) (err error) {
	si := &SagaInstance{}

	err = sm.repository.Load(result.Id, si)
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

func (sm *SagaManager) executeStep(si *SagaInstance) (err error) {
	for {
		if si.checkFinishState(sm.sagaDefinition.Len()) {
			break
		}

		step := sm.sagaDefinition.Get(si.StepIndex)

		if !si.checkStepHasHandler(step) {
			// skip if step does not have handler
			si.shiftIndex()
			continue
		}

		result, err := si.executeStepHandler(step)
		if err != nil {
			return err
		}

		sm.resultChannel <- result
		break
	}

	return
}
