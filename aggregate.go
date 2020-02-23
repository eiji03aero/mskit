package mskit

type Aggregate interface {
	Validate() []error
	Process(cmd interface{}) (Events, error)
	Apply(event interface{}) error
}

type BaseAggregate struct {
	Id string `json:"id"`
}

func (_ *BaseAggregate) Validate() []error {
	return nil
}
