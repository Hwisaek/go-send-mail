package send_grid

import (
	"github.com/Hwisaek/go-send-mail/model"
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		request model.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				request: model.Request{
					From: model.Mail{
						Name:  "test",
						Email: "test@hwisaek.com",
					},
					ToList: []model.Mail{
						{
							Name:  "Hwisaek",
							Email: "hwisaek@gmail.com",
						},
					},
					CcList:  nil,
					BccList: nil,
					Subject: "test",
					Content: "테스트입니다",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
