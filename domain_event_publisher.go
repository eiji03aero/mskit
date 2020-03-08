package mskit

type DomainEventPublisher interface {
	Publish(event interface{}) error
}

type StubDomainEventPublisher struct{}

func (s *StubDomainEventPublisher) Publish(event interface{}) error {
	return nil
}
