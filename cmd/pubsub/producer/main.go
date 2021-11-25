package main

import (
	"fmt"
	"log"
	"time"

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

	newUsersCount := 1
	for {
		newUserMessage := &protoMessages.NewUser{
			User: &protoMessages.User{
				Id:    uint32(newUsersCount),
				Email: fmt.Sprintf("lucas-%d@gmail.com", uint32(newUsersCount)),
			},
		}

		newUsersCount += 1

		newUserPayload, err := proto.Marshal(newUserMessage)
		if err != nil {
			log.Fatal(err)
		}

		err = amqpChannel.Publish(
			newlyRegisteredUsersExchange,
			"",
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         newUserPayload,
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Millisecond * 500)
	}
}
