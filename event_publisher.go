package mskit

// EventPublisher defines interface for publisher
type EventPublisher interface {
	Publish(event interface{}) error
}

// StubEventPublisher is a stub for the case it have to give but not needed
type StubEventPublisher struct{}

// Publish is stub implementation
func (s *StubEventPublisher) Publish(event interface{}) error {
	return nil
}
