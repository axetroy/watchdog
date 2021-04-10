package watchdog

import (
	"strings"
	"sync"
	"time"
)

// 限流的报警器，不至于过于频繁的报警
// 定义一个报警策略
type Alarm struct {
	mux               sync.Mutex
	name              string        // 唯一 ID
	interval          time.Duration // 触发两次报警的间隔时间
	lastTriggerAt     *time.Time    // 上一次触发时间
	triggerTimesToDay uint          // 今日已触发次数
	maxTimesForDay    uint          // 每天最多触发次数，0 无限次数
	maxTimesForHour   uint          // 每小时最多触发次数，0 无限次数
	Rest              []string      // 休息时段，这个时段内，不会发送通知，格式是 hh:mm:ss~hh:mm:ss，例如 02:00:00~09:00:00
}

type AlertOptions struct {
	Name            string        // 唯一 ID
	Interval        time.Duration // 触发两次报警的间隔时间
	MaxTimesForDay  uint          // 每日最多触发上限
	MaxTimesForHour uint          // 每小时最多触发上限
}

func NewAlarm(options AlertOptions) *Alarm {
	return &Alarm{
		interval:        options.Interval,
		maxTimesForDay:  options.MaxTimesForDay,
		maxTimesForHour: options.MaxTimesForHour,
		Rest:            make([]string, 0),
	}
}

func (a *Alarm) Tick() (shouldTrigger bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	defer func() {
		if shouldTrigger {
			t := time.Now()

			// 如果是同一天
			if a.lastTriggerAt != nil && t.Day() == a.lastTriggerAt.Day() {
				a.triggerTimesToDay = a.triggerTimesToDay + 1
			} else {
				a.triggerTimesToDay = 1
			}

			a.lastTriggerAt = &t
		}
	}()

	now := time.Now()

	// 如果超出了当前报警的量
	if a.maxTimesForDay > 0 && a.lastTriggerAt != nil {
		if now.Day() == a.lastTriggerAt.Day() && a.triggerTimesToDay >= a.maxTimesForDay {
			shouldTrigger = false
			return
		}
	}

	// 如果超出了这个小时内的量
	if a.maxTimesForHour > 0 {
		history := Store.GetItem(a.name)
		if history != nil {
			currentHour := now.Hour()

			errorHistory := filter(*history, func(ss ServiceStatus) bool {
				t, err := time.Parse(time.RFC3339, ss.UpdatedAt)

				if err != nil {
					return false
				}

				return ss.Error != "" && t.Hour() == currentHour
			})

			if len(errorHistory) >= int(a.maxTimesForHour) {
				shouldTrigger = false
				return
			}
		}
	}

	// 查看是否在休息时段外
	if a.Rest != nil {
		for _, ranges := range a.Rest {
			arr := strings.Split(ranges, "~")
			start, err := time.Parse(time.Stamp, "Jan _2 "+arr[0])

			if err != nil {
				break
			}

			end, err := time.Parse(time.Stamp, "Jan _2 "+arr[1])

			if err != nil {
				break
			}

			if now.After(start) && now.Before(end) {
				// do nothing
				shouldTrigger = true
			} else {
				shouldTrigger = false
				return
			}
		}
	}

	if a.lastTriggerAt == nil {
		shouldTrigger = true
		return
	}

	// 如果还没有到下一次的报警时间
	if now.After(a.lastTriggerAt.Add(a.interval)) {
		shouldTrigger = true
		return
	}

	return false
}

func filter(vs []ServiceStatus, f func(ServiceStatus) bool) []ServiceStatus {
	filtered := make([]ServiceStatus, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
