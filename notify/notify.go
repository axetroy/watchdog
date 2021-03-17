package notify

import "github.com/axetroy/watchdog"

type Notifier interface {
	Push(r watchdog.RunnerResult) error
}
