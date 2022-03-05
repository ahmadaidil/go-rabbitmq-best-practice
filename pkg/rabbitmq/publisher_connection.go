package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//Publish publishes a request to the amqp queue
func (c *Connection) Publish(m Message) error {
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		ContentType:   m.ContentType,
		CorrelationId: m.CorrelationID,
		Body:          m.Body,
		ReplyTo:       m.ReplyTo,
	}
	if err := c.channel.Publish("", m.Queue, false, false, p); err != nil {
		log.Printf("error in publishing: %s", err.Error())
		return fmt.Errorf("error in publishing: %s", err.Error())
	}

	log.Printf("Publish message to queue %q: %s", m.Queue, m.Body)
	return nil
}
