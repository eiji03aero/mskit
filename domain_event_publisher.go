package mskit

// DomainEventPublisher defines interface for publisher
type DomainEventPublisher interface {
	Publish(event interface{}) error
}

// StubDomainEventPublisher is a stub for the case it have to give but not needed
type StubDomainEventPublisher struct{}

// Publish is stub implementation
func (s *StubDomainEventPublisher) Publish(event interface{}) error {
	return nil
}
