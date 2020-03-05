package mskit

type DomainEventPublisher interface {
	Publish(event interface{}) error
}
