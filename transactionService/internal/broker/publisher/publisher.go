package publisher

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	event2 "transactionService/internal/broker/event"
)

func Publish(rabbitMqConnect *amqp.Connection, event event2.Event, queueName string) error {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in Publish", err)
		}
	}()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(event); err != nil {
		return fmt.Errorf("error encode event %w", err)
	}

	ch, err := rabbitMqConnect.Channel()
	if err != nil {
		return fmt.Errorf("failed to initialize rabbit mq channel: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	err = ch.PublishWithContext(context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        buf.Bytes(),
		})
	if err != nil {
		return fmt.Errorf("error publish event %w", err)
	}
	return nil
}
