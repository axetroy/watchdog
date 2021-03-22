package notify

import (
	"github.com/axetroy/watchdog"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

func Wechat(content string, reporter watchdog.Reporter) (err error) {
	type Payload struct {
		AppToken string `json:"app_token" mapstructure:"app_token"`
	}

	payload := Payload{}

	if err = mapstructure.Decode(reporter.Payload, &payload); err != nil {
		return errors.WithMessage(err, "微信配置不正确")
	}

	msg := model.
		NewMessage(payload.AppToken).
		SetContent(content).SetSummary(content)
	if len(reporter.Target) > 1 {
		msg = msg.AddUId(reporter.Target[0], reporter.Target[1:]...)
	} else if len(reporter.Target) == 1 {
		msg = msg.AddUId(reporter.Target[0])
	} else {
		return nil
	}

	if _, err = wxpusher.SendMessage(msg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
