package usecase

import "github.com/andres06-hub/events-example/internal/domain"

type EventUseCase struct {
	repo domain.EventRepository
}

func NewEventUseCase(r domain.EventRepository) *EventUseCase {
	return &EventUseCase{repo: r}
}

func (uc *EventUseCase) PublishEvent(event domain.Event) error {
	return uc.repo.Publish(event)
}

func (uc *EventUseCase) ConsumeEvents(handler func(domain.Event)) error {
	return uc.repo.Consume(handler)
}
