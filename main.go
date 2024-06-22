package main

import (
	"erp-mailer/event"
	"log"
	"os"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	ToAddress   string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

func main() {
	rabbit, err := ConnectRabbitMQ()

	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
		os.Exit(1)
	}
	defer rabbit.Close()

	consumer, err := event.NewConsumer(rabbit)
	if err != nil {
		panic(err)
	}
	err = consumer.Listen()
	if err != nil {
		panic(err)
	}
}
