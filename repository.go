package mskit

type Repository interface {
	Save(event *Event) error
	// Load
}
