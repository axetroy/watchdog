package scheduling

import (
	"fmt"
	"log"
	"time"

	"github.com/axetroy/watchdog"
	"github.com/axetroy/watchdog/internal/notify"
	"github.com/axetroy/watchdog/internal/socket"
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
		defer ticker.Stop()
		for range ticker.C {
			service := s.job.GetService()
			t1 := time.Now()
			err := s.job.Do()
			duration := time.Since(t1)

			data := socket.Data{
				Event: socket.EventUpdate,
			}

			serviceStatus := watchdog.ServiceStatus{
				Name:      service.Name,
				UpdatedAt: time.Now().Format(time.RFC3339),
				Duration:  duration,
			}

			if err != nil {
				log.Printf(`「%s」: %s`, service.Name, err.Error())
				serviceStatus.Error = err.Error()

				// 如果报错的话，检查是否应该上报错误
				if s.alarm.Tick() {
					// 开始推送
					pusher := notify.NewNotifier(service.Report)

					msg := fmt.Sprintf(`「watchdog」服务 '%s' 不可用, %s`, service.Name, err.Error())

					if pushErr := pusher.Push(msg); pushErr != nil {
						// write to log
						log.Println(pushErr)
					}
				}
			}

			data.Payload = serviceStatus
			socket.Pool.Broadcast(data)

			watchdog.Store.SetItem(serviceStatus.Name, []watchdog.ServiceStatus{serviceStatus})
		}
		ch <- 1
	}()
	<-ch
}
