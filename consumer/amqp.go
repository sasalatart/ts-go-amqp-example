package main

import (
	"os"

	"github.com/streadway/amqp"
)

const queueName = "ts-go-amqp-example"

func initConnection() (*amqp.Connection, error) {
	amqpURI, ok := os.LookupEnv("AMQP_URI")
	if !ok {
		amqpURI = "amqp://guest:guest@localhost:5672"
	}

	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func consumeQueue(ch *amqp.Channel, queue *amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
}
