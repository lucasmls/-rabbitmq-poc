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

	for {
		newUserMessage := &protoMessages.NewUser{
			User: &protoMessages.User{
				Id:    rand.Uint32() % 1000,
				Email: fmt.Sprintf("lucas-%d@gmail.com", rand.Intn(100)),
			},
		}

		newUserPayload, err := proto.Marshal(newUserMessage)
		if err != nil {
			log.Fatal(err)
		}

		err = amqpChannel.Publish(
			"",
			newUsersQueue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        newUserPayload,
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Millisecond * 500)
	}
}
