package notify

import (
	"net/smtp"

	"github.com/axetroy/watchdog"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func EmailSMTP(content string, reporter watchdog.Reporter) error {
	type Payload struct {
		Addr     string `json:"addr" mapstructure:"addr"`         // example: mail.example.com:25
		From     string `json:"from" mapstructure:"from"`         // example: sender@example.com
		Username string `json:"username" mapstructure:"username"` // example: user@example.com
		Password string `json:"password" mapstructure:"password"` // example: password
		Host     string `json:"host" mapstructure:"host"`         // example: mail.example.com
	}

	payload := Payload{}

	if err := mapstructure.Decode(reporter.Payload, &payload); err != nil {
		return errors.WithMessage(err, "邮箱服务器配置不正确")
	}

	auth := smtp.PlainAuth("", payload.Username, payload.Password, payload.Host)
	err := smtp.SendMail(payload.Addr, auth, payload.From, reporter.Target, []byte(content))

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
