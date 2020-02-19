package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

type amqpMessage struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func randomSleep() {
	time.Sleep(time.Duration(5+rand.Intn(10)) * time.Second)
}

func processDelivery(delivery amqp.Delivery, workerID int) {
	message := amqpMessage{}

	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		log.Printf("[w#%d] Error unmarshalling message: %v\n", workerID, err)
		return
	}

	id := message.ID
	log.Printf("[w#%d] Delivery %s received. Processing...", workerID, id)
	randomSleep()
	log.Printf("[w#%d] Delivery %s had message %q.\n", workerID, id, message.Message)
	delivery.Ack(false)
}

func work(messages <-chan amqp.Delivery, workerID int) {
	log.Printf("Starting worker #%d...", workerID)
	for delivery := range messages {
		processDelivery(delivery, workerID)
	}
}
