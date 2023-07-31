package consumer

import (
	"bytes"
	"encoding/gob"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	event2 "transactionService/internal/broker/event"
	"transactionService/internal/broker/listener"
)

func Consume(
	rabbitMqConnect *amqp.Connection,
	event event2.Event,
	queueName string,
	listener *listener.Listener,
) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println("panic in Consume", err)
		}
	}()

	ch, err := rabbitMqConnect.Channel()
	if err != nil {
		fmt.Errorf("failed to initialize rabbit mq channel: %w", err)
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
	if err != nil {
		logrus.Fatalf("Error QueueDeclare %s", err.Error())
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logrus.Fatalf("Error consume event %s", err.Error())
	}

	go func() {
		for d := range msgs {
			dec := gob.NewDecoder(bytes.NewBuffer(d.Body))
			if err = dec.Decode(&event); err != nil {
				logrus.Fatalf("error encode event %s", err.Error())
			}
			logrus.Printf("Received a message: %s", event)
			logrus.Info(listener.Handle(event))
		}
	}()

}
