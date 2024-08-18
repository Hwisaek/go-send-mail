package send_mail

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"os"
)

func sendgridSendMail(from Mail, toList, ccList, bccList []Mail, subject, content string) error {
	m := mail.NewV3Mail()

	m.SetFrom(mail.NewEmail(from.Name, from.Email))

	personalization := mail.NewPersonalization()
	personalization.AddTos(lo.Map(toList, func(item Mail, index int) *mail.Email {
		return mail.NewEmail(item.Name, item.Email)
	})...)
	personalization.AddCCs(lo.Map(ccList, func(item Mail, index int) *mail.Email {
		return mail.NewEmail(item.Name, item.Email)
	})...)
	personalization.AddBCCs(lo.Map(bccList, func(item Mail, index int) *mail.Email {
		return mail.NewEmail(item.Name, item.Email)
	})...)
	m.AddPersonalizations(personalization)

	m.Subject = subject

	m.AddContent(mail.NewContent("text/html", content))

	trackingSettings := mail.NewTrackingSettings()
	clickTrackingSetting := mail.NewClickTrackingSetting()
	clickTrackingSetting.SetEnable(true)
	clickTrackingSetting.SetEnableText(true)
	trackingSettings.SetClickTracking(clickTrackingSetting)
	openTrackingSetting := mail.NewOpenTrackingSetting()
	openTrackingSetting.SetEnable(true)
	openTrackingSetting.SetSubstitutionTag("%open-track%")
	trackingSettings.SetOpenTracking(openTrackingSetting)
	subscriptionTrackingSetting := mail.NewSubscriptionTrackingSetting()
	subscriptionTrackingSetting.SetEnable(true)
	trackingSettings.SetSubscriptionTracking(subscriptionTrackingSetting)
	m.SetTrackingSettings(trackingSettings)

	sendGridRequest := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	sendGridRequest.Method = http.MethodPost
	sendGridRequest.Body = mail.GetRequestBody(m)

	response, err := sendgrid.API(sendGridRequest)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("%s", response.Body)
	}

	return nil
}
