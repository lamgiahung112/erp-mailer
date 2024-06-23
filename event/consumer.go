package event

import (
	"crypto/tls"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type MailRequestPayload struct {
	MailType string `json:"name"`
	Data     any    `json:"data"`
}

var dialer *gomail.Dialer

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
	setupDialer()
	log.Println("Listening for mail requests!")
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

func setupDialer() {
	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	dialer = gomail.NewDialer(host, port, username, password)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	log.Println("Mail dialer is set up!")
}

func handlePayload(p *MailRequestPayload) {
	switch MailType(p.MailType) {
	case LoginOTP:
		sendLoginOtpEmail(&p.Data)
	case VerifyAccount:
		log.Println("VerifyAccount")
	default:
		log.Println("Not ok")
	}
}
