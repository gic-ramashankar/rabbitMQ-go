package service

import (
	"demo/model"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

func InsertInQueue(body model.Request) (string, error) {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Error: ", err)
		return "", errors.New("Failed Initializing Broker Connection")

	}

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	a, err := json.Marshal(body)
	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			//type -> json
			ContentType: "json/application",
			Body:        []byte(a),
		},
	)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return "Successfully Published Message to Queue", nil
}
