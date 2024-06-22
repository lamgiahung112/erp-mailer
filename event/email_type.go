package event

type MailType string

func (e MailType) String() string {
	return string(e)
}

const (
	LoginOTP      = MailType("login_otp")
	VerifyAccount = MailType("verify_account")
)
