package mskit

// Aggregate defines interface for aggregate
type Aggregate interface {
	Validate() []error
	Process(cmd interface{}) (Events, error)
	Apply(event interface{}) error
}

// BaseAggregate is a struct to express aggregate
type BaseAggregate struct {
	Id string `json:"id"`
}

// Validate verifies if aggregate is valid. Could be run in different spots
func (_ *BaseAggregate) Validate() []error {
	return nil
}
