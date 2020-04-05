package mskit

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
func (r *EventRepository) Save(aggregate Aggregate, event Event) error {
	err := aggregate.Apply(event.Data)
	if err != nil {
		return err
	}

	errors := aggregate.Validate()
	if len(errors) > 0 {
		return errors[0]
	}

	err = r.eventStore.Save(event)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteCommand executes, saves and publishes data
func (r *EventRepository) ExecuteCommand(aggregate Aggregate, cmd interface{}) error {
	events, err := aggregate.Process(cmd)
	if err != nil {
		return err
	}

	for _, e := range events {
		err = r.Save(aggregate, e)
		if err != nil {
			return err
		}

		err = r.publisher.Publish(e.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

// Load loads up data for aggregate
func (r *EventRepository) Load(id string, aggregate Aggregate) error {
	err := r.eventStore.Load(id, aggregate)
	if err != nil {
		return err
	}

	return nil
}
