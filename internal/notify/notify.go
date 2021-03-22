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

type Notify struct {
	reporter []watchdog.Reporter
	handler  Handler
}

type HandlerFn = func(message string, reporter watchdog.Reporter) error

type Handler = map[string]HandlerFn

func NewNotifier(ns []watchdog.Reporter) Notifier {
	notify := Notify{
		reporter: ns,
		handler:  map[string]HandlerFn{},
	}

	notify.use("wechat", Wechat)
	notify.use("webhook", Webhook)

	return notify
}

// 注册处理器
func (n *Notify) use(protocol string, handler HandlerFn) {
	n.handler[protocol] = handler
}

func (n Notify) Push(content string) error {
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

func (n Notify) PushEmail(msg string) error {
	return nil
}
