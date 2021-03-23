package watchdog

import (
	"context"
	"time"

	"github.com/axetroy/watchdog/internal/protocol"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

// 如果返回 error，说明服务异常
func processSingleService(s Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	err := protocol.Ping(s.Protocol, s.Addr, ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewRunnerJob(s Service) RunnerJob {
	return RunnerJob{
		service: s,
	}
}

type RunnerJob struct {
	service Service
}

func (r RunnerJob) Name() string {
	return r.service.Name
}

func (r RunnerJob) GetService() Service {
	return r.service
}

func (r RunnerJob) Do() error {
	err := processSingleService(r.service)

	if err != nil {
		color.Red.Printf("[%s]: 异常\n", r.service.Name)
	} else {
		color.Green.Printf("[%s]: 正常\n", r.service.Name)
	}

	return err
}
