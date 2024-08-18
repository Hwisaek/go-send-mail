package send_mail

import (
	"github.com/Hwisaek/go-send-mail/model"
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		from    model.Mail
		toList  []model.Mail
		ccList  []model.Mail
		bccList []model.Mail
		subject string
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				from: model.Mail{
					Name:  "test",
					Email: "test@hwisaek.com",
				},
				toList: []model.Mail{
					{
						Name:  "Hwisaek",
						Email: "hwisaek@gmail.com",
					},
				},
				subject: "test",
				content: "테스트입니다",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.from, tt.args.toList, tt.args.ccList, tt.args.bccList, tt.args.subject, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
