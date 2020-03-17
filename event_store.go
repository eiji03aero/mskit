package mskit

// EventStore is the interface to give the basic usage on underlying data store
type EventStore interface {
	Save(event Event) error
	Load(id string, aggregate Aggregate) error
}
