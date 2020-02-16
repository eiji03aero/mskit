package mskit

import (
	"github.com/eiji03aero/mskit/utils"
)

type Event struct {
	ID            string
	Type          string
	AggregateType string
	AggregateID   string
	Data          interface{}
}

func NewEvent(
	aggregateID string,
	aggregate interface{},
	event interface{},
) *Event {
	_, aggregateType := utils.GetTypeName(aggregate)
	_, eventType := utils.GetTypeName(event)

	return &Event{
		Type:          eventType,
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		Data:          event,
	}
}
