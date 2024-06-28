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
	msg.SetHeader("From", os.Getenv("MAIL_USERNAME"))
	msg.SetHeader("To", castData.ToAddress)
	msg.SetHeader("Subject", castData.Title)
	msg.SetHeader("MIME-Version", "1.0")
	msg.SetHeader("Content-Type", "text/html; charset=utf-8")
	msg.SetBody("text/html", fmt.Sprintf(`
		<div>
			Hi, %s,<br>
			Your OTP is: %s
		</div>`,
		castData.Username, castData.Otp,
	))
	err := dialer.DialAndSend(msg)
	if err != nil {
		fmt.Println("dialer.DialAndSend err:", err)
	}
	log.Println("Sending login otp email success")
}
