package watchdog

import (
	"context"
	"time"

	"github.com/axetroy/watchdog/internal/protocol"
	"github.com/pkg/errors"
)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	err := protocol.Ping(r.service.Protocol, r.service.Addr, ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
