package notify

import (
	"log"
	"sync"

	"github.com/axetroy/watchdog"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

type Notifier interface {
	Push(msg string) error
}

type Notify struct {
	notification []watchdog.Notification
}

func NewNotifier(ns []watchdog.Notification) Notifier {
	return Notify{
		notification: ns,
	}
}

func (n Notify) Push(content string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(n.notification))

	for _, notification := range n.notification {
		go func(notification watchdog.Notification) {
			var (
				err error
			)

			defer func() {
				if err != nil {
					log.Printf("%+v\n", err)
				}
			}()

			defer wg.Done()

			switch notification.Protocol {
			case "wechat":
				type Payload struct {
					AppToken string `json:"app_token" mapstructure:"app_token"`
				}

				payload := Payload{}

				if err = mapstructure.Decode(notification.Payload, &payload); err != nil {
					err = errors.WithMessage(err, "微信配置不正确")
					return
				}

				msg := model.
					NewMessage(payload.AppToken).
					SetContent(content).SetSummary(content)
				if len(notification.Target) > 1 {
					msg = msg.AddUId(notification.Target[0], notification.Target[1:]...)
				} else if len(notification.Target) == 1 {
					msg = msg.AddUId(notification.Target[0])
				} else {
					return
				}

				_, err = wxpusher.SendMessage(msg)
				if err != nil {
					return
				}

			default:
				err = errors.Errorf("invlid protocol '%s'", notification.Protocol)
				return
			}
		}(notification)
	}

	wg.Wait()
	return nil
}

func (n Notify) PushEmail(msg string) error {
	return nil
}
