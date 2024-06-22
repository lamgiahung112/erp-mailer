package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	var connection *amqp.Connection

	c, err := amqp.Dial("amqp://guest:guest@rabbitmq")

	if err != nil {
		return nil, err
	}
	connection = c
	return connection, nil
}
