package handler

import (
	"log"

	rbmq "github.com/ahmadaidil/go-rabbitmq-best-practice/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func ConsumerMessageHandler(c rbmq.Connection, q string, deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		m := rbmq.Message{
			Queue:         q,
			Body:          d.Body,
			ContentType:   d.ContentType,
			Priority:      d.Priority,
			CorrelationID: d.CorrelationId,
		}
		//handle the message
		log.Printf("Got message from queue %q: %s", m.Queue, m.Body)
		d.Ack(false)
	}
}
