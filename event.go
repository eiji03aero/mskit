package mskit

import (
	"github.com/eiji03aero/mskit/utils"
)

// Event is struct to express domain event
type Event struct {
	Id            string
	Type          string
	AggregateType string
	AggregateId   string
	Data          interface{}
}

// Events is a type to express slice of Event
type Events = []Event

// NewEvent is utility function to create Event struct
func NewEvent(
	aggregateId string,
	aggregate interface{},
	event interface{},
) Event {
	aggregateType := utils.GetTypeName(aggregate)
	eventType := utils.GetTypeName(event)

	return Event{
		Type:          eventType,
		AggregateId:   aggregateId,
		AggregateType: aggregateType,
		Data:          event,
	}
}

// NewEventsSingle is utility to create a single event in slice
func NewEventsSingle(
	aggregateId string,
	aggregate interface{},
	event interface{},
) Events {
	return Events{
		NewEvent(
			aggregateId,
			aggregate,
			event,
		),
	}
}
