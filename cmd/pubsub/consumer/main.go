package main

import (
	"fmt"
	"log"

	protoMessages "github.com/lucasmls/rabbitmq-poc/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

const (
	newlyRegisteredUsersExchange string = "new_users"
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

	err = amqpChannel.ExchangeDeclare(
		newlyRegisteredUsersExchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	temporaryNewUsersQueue, err := amqpChannel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = amqpChannel.QueueBind(
		temporaryNewUsersQueue.Name,
		"",
		newlyRegisteredUsersExchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	messages, err := amqpChannel.Consume(
		temporaryNewUsersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for message := range messages {
		newUserMessage := &protoMessages.NewUser{}
		err := proto.Unmarshal(message.Body, newUserMessage)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(newUserMessage)
	}
}
