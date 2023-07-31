package tests

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"testing"
	"transactionService/internal/broker/event"
	"transactionService/internal/broker/publisher"
)

func TestPublisher(t *testing.T) {
	rabbitMqConnect, _ := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", "broker", "broker", "localhost", "5672"))

	err := publisher.Publish(rabbitMqConnect, event.Event{}, "test-queue")
	if err != nil {
		t.Errorf("fall  %s", err)
	}
}
