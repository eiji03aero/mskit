package mskit

// SagaDefinition is a struct to express the abstracted definition of saga
type SagaDefinition struct {
	steps []SagaStep
}

// NewSagaDefinition returns new SagaDefinition
func NewSagaDefinition() *SagaDefinition {
	return &SagaDefinition{}
}

func (sd *SagaDefinition) addStep() {
	sd.steps = append(sd.steps, SagaStep{})
}

func (sd *SagaDefinition) Len() int {
	return len(sd.steps)
}

func (sd *SagaDefinition) Get(index int) (ss SagaStep) {
	return sd.steps[index]
}

// SagaDefinitionBuilder is a struct that is used while building Saga
type SagaDefinitionBuilder struct {
	sagaDefinition *SagaDefinition
}

// NewSagaDefinitionBuilder returns new SagaDefinitionBuilder with necessary initialization
func NewSagaDefinitionBuilder() *SagaDefinitionBuilder {
	return &SagaDefinitionBuilder{
		sagaDefinition: NewSagaDefinition(),
	}
}

// Step appends new step to SagaDefinition
func (b *SagaDefinitionBuilder) Step(opts ...interface{}) *SagaDefinitionBuilder {
	b.sagaDefinition.addStep()
	step := b.getCurrentStep()

	step.Configure(opts)

	return b
}

// Build wraps up building SagaDefinition
func (b *SagaDefinitionBuilder) Build() (*SagaDefinition, error) {
	for _, step := range b.sagaDefinition.steps {
		err := step.Validate()
		if err != nil {
			return nil, err
		}
	}

	return b.sagaDefinition, nil
}

func (b *SagaDefinitionBuilder) getCurrentStep() SagaStep {
	steps := b.sagaDefinition.steps
	return steps[len(steps)-1]
}
