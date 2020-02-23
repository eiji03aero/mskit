package mskit

type Repository struct {
	eventStore EventStore
}

func NewRepository(eventStore EventStore) *Repository {
	return &Repository{
		eventStore: eventStore,
	}
}

func (r *Repository) Save(aggregate Aggregate, event *Event) error {
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

func (r *Repository) Load(id string, aggregate Aggregate) error {
	err := r.eventStore.Load(id, aggregate)
	if err != nil {
		return err
	}

	return nil
}
