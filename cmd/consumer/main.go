package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	ampqConnection, err := amqp.Dial("amqp:guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	defer ampqConnection.Close()

	amqpChannel, err := ampqConnection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	newUsersQueue, err := amqpChannel.QueueDeclare(
		"new-users",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	messages, err := amqpChannel.Consume(
		newUsersQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for message := range messages {
		fmt.Printf("Received a new message, content: %s \n", message.Body)
	}
}
