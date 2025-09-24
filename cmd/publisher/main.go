package main

import (
	"os"

	"github.com/andres06-hub/events-example/internal/domain"
	"github.com/andres06-hub/events-example/internal/infrastructure/rabbitmq"
	"github.com/andres06-hub/events-example/internal/usecase"
)

func main() {
	rabbirConn := os.Getenv("RABBITMQ_CONN")
	if rabbirConn == "" {
		rabbirConn = "amqp://guest:guest@localhost:5672/"
	}
	conn := rabbitmq.NewRabbitMQConnection(rabbirConn)
	defer conn.Close()

	repo := rabbitmq.NewRabbitPublisher(conn, "events_queue")
	uc := usecase.NewEventUseCase(repo)

	event := domain.Event{
		ID:      "1",
		Type:    "UserCreated",
		Payload: "New user successfully created",
	}

	uc.PublishEvent(event)
}
