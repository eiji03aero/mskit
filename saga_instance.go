package mskit

import (
	"fmt"

	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
)

// SagaInstanceState defines state of SagaInstance
type SagaInstanceState int

const (
	SagaInstanceState_Unknown SagaInstanceState = iota
	SagaInstanceState_Processing
	SagaInstanceState_Aborting
	SagaInstanceState_Aborted
	SagaInstanceState_Done
)

// SagaInstace is a struct to express an instance of saga
type SagaInstance struct {
	Id        string            `json:"id"`
	StepIndex int               `json:"step_index"`
	State     SagaInstanceState `json:"state"`
	Data      interface{}       `json:"data"`
}

func NewSagaInstance() (si *SagaInstance, err error) {
	id, err := utils.UUID()
	if err != nil {
		return
	}

	si = &SagaInstance{
		Id:        id,
		StepIndex: 0,
		State:     SagaInstanceState_Processing,
	}

	return
}

func (si *SagaInstance) processResult(result *SagaStepResult) (err error) {
	if result.Error != nil && si.State == SagaInstanceState_Processing {
		logger.PrintFuncCall(si.processResult, logger.RedString("aborting"), si, result)
		si.State = SagaInstanceState_Aborting
		// Need to return, since current step might have compensation
		return
	}

	err = si.shiftIndex()

	return
}

func (si *SagaInstance) shiftIndex() (err error) {
	switch si.State {
	case SagaInstanceState_Processing:
		si.StepIndex++
	case SagaInstanceState_Aborting:
		si.StepIndex--
	default:
		return fmt.Errorf("inproper state for saga instance. id=%s state=%d", si.Id, si.State)
	}
	return
}

func (si *SagaInstance) checkFinishState(lenSteps int) bool {
	i := si.StepIndex
	switch {
	case i < 0:
		si.State = SagaInstanceState_Aborted
		return true
	case i >= lenSteps:
		si.State = SagaInstanceState_Processing
		return true
	default:
		return false
	}
}

func (si *SagaInstance) checkStepHasHandler(step *SagaStep) bool {
	switch si.State {
	case SagaInstanceState_Processing:
		return step.executeHandler != nil
	case SagaInstanceState_Aborting:
		return step.compensationHandler != nil
	default:
		return false
	}
}

func (si *SagaInstance) executeStepHandler(step *SagaStep) (err error) {
	switch si.State {
	case SagaInstanceState_Processing:
		funcName := utils.GetFunctionNameParent(step.executeHandler)
		logger.PrintFuncCall(si.executeStepHandler, logger.RedString("executing"), funcName, si.Data)
		err = step.executeHandler(si.Data)
	case SagaInstanceState_Aborting:
		funcName := utils.GetFunctionNameParent(step.compensationHandler)
		logger.PrintFuncCall(si.executeStepHandler, logger.RedString("compensating"), funcName, si.Data)
		err = step.compensationHandler(si.Data)
	default:
		err = fmt.Errorf("inproper state for saga instance. id=%s state=%d", si.Id, si.State)
	}
	return
}
