package main

import (
	"log"

	"github.com/ahmadaidil/go-rabbitmq-best-practice/pkg/handler"
	rbmq "github.com/ahmadaidil/go-rabbitmq-best-practice/pkg/rabbitmq"
)

func main() {
	forever := make(chan bool)
	conn := rbmq.NewConnection(
		"my-consumer",               //name
		"my-exchange",               //exchange
		[]string{"user", "product"}, //queues
	)
	if err := conn.Connect(); err != nil {
		panic(err)
	}
	if err := conn.BindQueue(); err != nil {
		panic(err)
	}
	deliveries, err := conn.Consume()
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully connected to RabbitMQ server as consumer!")
	for q, d := range deliveries {
		go conn.HandleConsumedDeliveries(q, d, handler.ConsumerMessageHandler)
	}
	<-forever // keep alive
}
