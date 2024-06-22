package event

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type MailRequestPayload struct {
	MailType string `json:"name"`
	Data     string `json:"data"`
}

func NewConsumer(conn *amqp.Connection) (*Consumer, error) {
	consumer := &Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (consumer *Consumer) Listen() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	q, err := declareQueue(channel)
	if err != nil {
		return err
	}

	err = channel.QueueBind(q.Name, q.Name, "mail_topic", false, nil)
	if err != nil {
		return err
	}
	messages, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	log.Println("Listening for mail requests")
	forever := make(chan bool)
	go func() {
		for d := range messages {
			var p MailRequestPayload
			_ = json.Unmarshal(d.Body, &p)
			go handlePayload(&p)
		}
	}()
	<-forever
	return nil
}

func handlePayload(p *MailRequestPayload) {
	switch MailType(p.MailType) {
	case LoginOTP:
		log.Println("OK")
	case VerifyAccount:
		log.Println("VerifyAccount")
	default:
		log.Println("Not ok")
	}
}
