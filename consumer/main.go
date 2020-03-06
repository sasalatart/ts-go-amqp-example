package main

import (
	"log"
	"runtime"
	"time"

	"github.com/streadway/amqp"
)

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}

func createWorkers(ch *amqp.Channel, queue *amqp.Queue) {
	failures := 0
	for i := 0; i < runtime.NumCPU(); {
		msgs, err := consumeChannel(ch, queue)
		workerID := i + 1

		if err != nil {
			failures++
			retryingAt := time.Second << failures
			log.Printf("Failed to consume from queue for worker %d: %v.\n", workerID, err)
			log.Printf("Retrying in %s...", time.Duration(retryingAt/time.Second)*time.Second)
			time.Sleep(retryingAt)
			continue
		}

		go work(msgs, workerID)
		i++
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
