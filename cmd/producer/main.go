package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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

	for {
		newUserPayload := fmt.Sprintf("lucas-%d@gmail.com", rand.Intn(100))

		err = amqpChannel.Publish(
			"",
			newUsersQueue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(newUserPayload),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Millisecond * 500)
	}
}
