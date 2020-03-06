package main

import (
	"log"
	"runtime"

	"github.com/streadway/amqp"
)

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}

func createWorkers(ch *amqp.Channel, queue *amqp.Queue) {
	for i := 0; i < runtime.NumCPU(); i++ {
		msgs, err := consumeQueue(ch, queue)

		if err != nil {
			log.Printf("Failed to consume from queue: %v. Retrying...\n", err)
			i--
		}

		go work(msgs, i+1)
	}
}

func main() {
	conn, err := initConnection()
	handleError("Failed to initialize connection", err)
	defer conn.Close()

	ch, err := conn.Channel()
	handleError("Failed to initialize channel", err)
	ch.Qos(1, 0, false)
	defer ch.Close()

	queue, err := declareQueue(ch)
	handleError("Failed to initialize queue", err)

	listener := make(chan bool)
	log.Printf("%d logical CPUs available. Concurrency set to that number.\n", runtime.NumCPU())
	log.Println("Listening for messages...")
	createWorkers(ch, &queue)
	<-listener
}
