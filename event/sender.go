package event

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

type LoginOtpEmailPayload struct {
	ToAddress string `json:"to_address"`
	Title     string `json:"title"`
	Otp       string `json:"otp"`
	Username  string `json:"username"`
}

func mapLoginOtpEmailData(data *any) *LoginOtpEmailPayload {
	cast_data, ok := (*data).(map[string]interface{})
	if !ok {
		return nil
	}
	return &LoginOtpEmailPayload{
		ToAddress: cast_data["to_address"].(string),
		Title:     cast_data["title"].(string),
		Otp:       cast_data["otp"].(string),
		Username:  cast_data["username"].(string),
	}
}

func sendLoginOtpEmail(data *any) {
	castData := mapLoginOtpEmailData(data)
	if castData == nil {
		return
	}
	msg := gomail.NewMessage(gomail.SetCharset("UTF-8"), gomail.SetEncoding("UTF-8"))
	msg.FormatAddress(os.Getenv("MAIL_USERNAME"), "ERP Mailer")
	msg.SetHeader("From", os.Getenv("MAIL_USERNAME"))
	msg.SetHeader("To", "lamgiahung112@gmail.com")
	msg.SetHeader("Subject", "Hello")
	msg.SetBody("text/html", fmt.Sprintf(`<h1>%v</h1>`, castData))
	err := dialer.DialAndSend(msg)
	if err != nil {
		fmt.Println("dialer.DialAndSend err:", err)
	}
	log.Println("Sending login otp email success")
}
