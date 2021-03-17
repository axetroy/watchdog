package watchdog

import (
	"context"
	"time"

	"github.com/axetroy/watchdog/protocol"
	"github.com/pkg/errors"
)

type RunnerResult struct {
	Duration uint          `json:"times"`  // 运行本次任务所需时间
	Errors   []RunnerError `json:"errors"` // 运行本次任务的错误信息
}

type RunnerError struct {
	Name       string `json:"name"`        // 服务名称
	Error      string `json:"error"`       // 错误信息
	RetryTimes uint   `json:"retry_times"` // 重试次数
}

func Run(config Config) (r RunnerResult) {
	t1 := time.Now()

	for _, s := range config.Service {
		err := processSingleService(s)

		if err != nil {
			if r.Errors == nil {
				r.Errors = make([]RunnerError, 0)
			}

			r.Errors = append(r.Errors, RunnerError{
				Name:       s.Name,
				Error:      err.Error(),
				RetryTimes: 0,
			})
		}
	}

	t2 := time.Now()

	r.Duration = uint(t2.Sub(t1).Seconds())

	return r
}

func processSingleService(s Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	err := protocol.Ping(s.Protocol, s.Addr, ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
