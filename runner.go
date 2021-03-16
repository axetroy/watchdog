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
	Error      error  `json:"error"`       // 错误信息
	RetryTimes uint   `json:"retry_times"` // 重试次数
}

func Run(config Config) error {
	for _, s := range config.Service {
		err := processSingleService(s)

		if err != nil {
		}
	}

	return nil
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
