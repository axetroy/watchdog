package scheduling

import (
	"time"
)

type Job interface {
	Name() string
	Do() error
}

func NewScheduling(interval time.Duration, j Job) Scheduling {
	return Scheduling{
		interval: interval,
		job:      j,
	}
}

type Scheduling struct {
	interval time.Duration
	job      Job
}

// 开始调度
func (s *Scheduling) Start() {
	var ch chan int
	//定时任务
	ticker := time.NewTicker(s.interval)
	go func() {
		for range ticker.C {
			_ = s.job.Do()
		}
		ch <- 1
	}()
	<-ch
}
