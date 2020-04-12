package mskit

import (
	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
)

// EventRepository manages Event
type EventRepository struct {
	eventStore EventStore
	publisher  EventPublisher
}

// NewEventRepository creates new struct
func NewEventRepository(eventStore EventStore, publisher EventPublisher) *EventRepository {
	return &EventRepository{
		eventStore: eventStore,
		publisher:  publisher,
	}
}

// Save saves Event
func (r *EventRepository) save(aggregate Aggregate, event Event) (err error) {
	logger.Println(
		logger.HiBlueString("Save event on aggregate"),
		logger.CyanString(event.AggregateType),
		event.Data,
	)

	err = aggregate.Apply(event.Data)
	if err != nil {
		return
	}

	errors := aggregate.Validate()
	if len(errors) > 0 {
		return errors[0]
	}

	err = r.eventStore.Save(event)
	if err != nil {
		return
	}

	return nil
}

// SaveEvents saves Event and publish
func (r *EventRepository) Save(aggregate Aggregate, events Events) (err error) {
	for _, e := range events {
		err = r.save(aggregate, e)
		if err != nil {
			return
		}

		err = r.publisher.Publish(e.Data)
		if err != nil {
			return
		}
	}

	return nil
}

// ExecuteCommand executes, saves and publishes data
func (r *EventRepository) ExecuteCommand(aggregate Aggregate, cmd interface{}) (err error) {
	aggregateName := utils.GetTypeName(aggregate)
	logger.Println(
		logger.HiBlueString("Execute command on aggregate"),
		logger.CyanString(aggregateName),
		cmd,
	)

	events, err := aggregate.Process(cmd)
	if err != nil {
		return
	}

	err = r.Save(aggregate, events)
	if err != nil {
		return
	}

	return
}

// Load loads up data for aggregate
func (r *EventRepository) Load(id string, aggregate Aggregate) (err error) {
	aggregateName := utils.GetTypeName(aggregate)
	logger.Println(
		logger.HiBlueString("Load aggregate"),
		logger.CyanString(aggregateName),
		aggregate,
	)

	err = r.eventStore.Load(id, aggregate)
	if err != nil {
		return
	}
	return
}
