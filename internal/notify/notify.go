package notify

import (
	"log"
	"sync"

	"github.com/axetroy/watchdog"
	"github.com/pkg/errors"
)

type Notifier interface {
	Push(msg string) error
}

type notifier struct {
	reporter []watchdog.Reporter
	handler  handler
}

type handlerFn = func(message string, reporter watchdog.Reporter) error

type handler = map[string]handlerFn

// 创建新的消息通知器
func NewNotifier(ns []watchdog.Reporter) Notifier {
	notify := notifier{
		reporter: ns,
		handler:  map[string]handlerFn{},
	}

	notify.use("wechat", Wechat)
	notify.use("webhook", Webhook)
	notify.use("email", Email)
	return notify
}

// 注册处理器
func (n *notifier) use(protocol string, handler handlerFn) {
	n.handler[protocol] = handler
}

// 推送消息
func (n notifier) Push(content string) error {
	wg := sync.WaitGroup{}
	wg.Add(len(n.reporter))

	for _, reporter := range n.reporter {
		go func(reporter watchdog.Reporter) {
			var (
				err error
			)

			defer func() {
				if err != nil {
					log.Printf("%+v\n", err)
				}
			}()

			defer wg.Done()

			if handler, ok := n.handler[reporter.Protocol]; ok {
				//do something here
				err = handler(content, reporter)
			} else {
				err = errors.Errorf("invlid protocol '%s'", reporter.Protocol)
			}
		}(reporter)
	}

	wg.Wait()
	return nil
}
