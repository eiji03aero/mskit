package mskit

type EventStore interface {
	Save(event *Event) error
	Load(id string, aggregate Aggregate) error
}
