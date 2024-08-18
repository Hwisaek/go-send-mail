package send_grid

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"os"
	"send-mail/model"
)

func SendMail(request model.Request) error {
	m := mail.NewV3Mail()

	m.SetFrom(mail.NewEmail(request.From.Name, request.From.Email))

	personalization := mail.NewPersonalization()
	personalization.AddTos(lo.Map(request.ToList, func(item model.Mail, index int) *mail.Email {
		return mail.NewEmail(item.Name, item.Email)
	})...)
	personalization.AddCCs(lo.Map(request.CcList, func(item model.Mail, index int) *mail.Email {
		return mail.NewEmail(item.Name, item.Email)
	})...)
	personalization.AddBCCs(lo.Map(request.BccList, func(item model.Mail, index int) *mail.Email {
		return mail.NewEmail(item.Name, item.Email)
	})...)
	m.AddPersonalizations(personalization)

	m.Subject = request.Subject

	m.AddContent(mail.NewContent("text/html", request.Content))

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
