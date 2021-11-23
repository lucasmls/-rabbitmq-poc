package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	protoMessages "github.com/lucasmls/rabbitmq-poc/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
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
		message.Ack(false)
		time.Sleep(time.Duration(rand.Uint32()) % 3 * time.Second)
	}
}
