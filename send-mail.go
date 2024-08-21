package sendmail

func SendMail(from Mail, toList, ccList, bccList []Mail, subject, content string, attachmentList []Attachment) error {
	// TODO sendgrid 말고도 다른 시스템도 추가 연동
	return sendgridSendMail(from, toList, ccList, bccList, subject, content, attachmentList)
}
