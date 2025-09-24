package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func NewRabbitMQConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	return conn
}
