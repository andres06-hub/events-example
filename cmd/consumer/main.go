package main

import (
	"fmt"
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

	uc.ConsumeEvents(func(event domain.Event) {
		fmt.Println("Event received:", event)
	})

	select {} // Keep the program running
}
