package scheduling

import (
	"fmt"
	"log"
	"time"

	"github.com/axetroy/watchdog"
	"github.com/axetroy/watchdog/internal/notify"
)

type Job interface {
	Name() string
	Do() error
	GetService() watchdog.Service
}

func NewScheduling(interval time.Duration, j Job) Scheduling {
	return Scheduling{
		interval: interval,
		job:      j,
		alarm:    watchdog.NewAlarm(time.Minute*10, 100),
	}
}

type Scheduling struct {
	interval time.Duration
	job      Job
	alarm    *watchdog.Alarm
}

// 开始调度
func (s *Scheduling) Start() {
	var ch chan int
	//定时任务
	ticker := time.NewTicker(s.interval)
	go func() {
		for range ticker.C {
			err := s.job.Do()

			if err != nil {
				// 如果报错的话，检查是否应该上报错误
				if s.alarm.ShouldTrigger() {
					// 开始推送
					log.Println("开始推送")
					service := s.job.GetService()
					pusher := notify.NewNotifier(service.Notifycation)

					msg := fmt.Sprintf(`「watchdog」服务 '%s' 不可用, %s`, service.Name, err.Error())

					if pushErr := pusher.Push(msg); pushErr != nil {
						// write to log
						log.Println(pushErr)
					}
					log.Println("推送完毕")
				}
			}
		}
		ch <- 1
	}()
	<-ch
}
