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
	err := make(chan error, 1)
	timeout := time.Second * 30

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	go func() {
		e := protocol.Ping(r.service.Protocol, r.service.Addr, r.service.Auth, ctx)

		if e != nil {
			e = errors.WithStack(e)
			err <- e
		} else {
			err <- nil
		}
	}()

	select {
	case e := <-err:
		return e
	case <-ctx.Done():
		return errors.New("Timeout")
	}
}
