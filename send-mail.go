package send_mail

import (
	"github.com/Hwisaek/go-send-mail/model"
	"github.com/Hwisaek/go-send-mail/sendgrid"
)

func SendMail(from model.Mail, toList, ccList, bccList []model.Mail, subject, content string) error {
	// TODO sendgrid 말고도 다른 시스템도 추가 연동
	return sendgrid.SendMail(from, toList, ccList, bccList, subject, content)
}
