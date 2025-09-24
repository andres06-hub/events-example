package rabbitmq

import (
	"github.com/andres06-hub/events-example/internal/domain"
	"github.com/streadway/amqp"
)

type RabbitPublisher struct {
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewRabbitPublisher(conn *amqp.Connection, queueName string) *RabbitPublisher {
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	return &RabbitPublisher{ch: ch, queue: q}
}

func (r *RabbitPublisher) Publish(event domain.Event) error {
	return r.ch.Publish(
		"", r.queue.Name, false, false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(event.Payload)},
	)
}

func (r *RabbitPublisher) Consume(handler func(domain.Event)) error {
	msgs, err := r.ch.Consume(r.queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(domain.Event{Payload: string(d.Body)})
		}
	}()
	return nil
}
