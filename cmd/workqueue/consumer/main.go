package main

import (
	"fmt"
	"log"
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

	err = amqpChannel.Qos(1, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	newUsersQueue, err := amqpChannel.QueueDeclare(
		"new_users",
		true,
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
		time.Sleep(time.Second * 3)
	}
}
