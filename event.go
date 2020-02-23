package mskit

import (
	"github.com/eiji03aero/mskit/utils"
)

type Event struct {
	Id            string
	Type          string
	AggregateType string
	AggregateId   string
	Data          interface{}
}

type Events = []*Event

func NewEvent(
	aggregateId string,
	aggregate interface{},
	event interface{},
) *Event {
	_, aggregateType := utils.GetTypeName(aggregate)
	_, eventType := utils.GetTypeName(event)

	return &Event{
		Type:          eventType,
		AggregateId:   aggregateId,
		AggregateType: aggregateType,
		Data:          event,
	}
}

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
