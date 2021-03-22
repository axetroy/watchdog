package watchdog

import (
	"context"
	"sync"
	"time"

	"github.com/axetroy/watchdog/internal/protocol"
	"github.com/gookit/color"
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

	wg := sync.WaitGroup{}

	var ch = make(chan int, 10) // 最大并发数量 10

	for _, s := range config.Service {
		wg.Add(1)

		go func(s Service) {
			err := processSingleService(s)

			defer wg.Done()

			ch <- 1

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

			<-ch
		}(s)
	}

	wg.Wait()
	close(ch)

	t2 := time.Now()

	r.Duration = uint(t2.Sub(t1).Seconds())

	return r
}

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
