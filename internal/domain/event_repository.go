package domain

type EventRepository interface {
	Publish(event Event) error
	Consume(handler func(Event)) error
}
