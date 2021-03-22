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
			err := s.job.Do()

			data := socket.Data{
				Event: "update",
			}

			payload := map[string]string{
				"name":       service.Name,
				"updated_at": time.Now().Format(time.RFC3339),
			}

			if err != nil {
				payload["error"] = err.Error()

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

			data.Payload = payload
			socket.Pool.Broadcast(data)
		}
		ch <- 1
	}()
	<-ch
}
