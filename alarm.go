package watchdog

import (
	"sync"
	"time"
)

// 限流的报警器，不至于过于频繁的报警
// 定义一个报警策略
type Alarm struct {
	mux               sync.Mutex
	interval          time.Duration // 触发两次报警的间隔时间
	lastTriggerAt     *time.Time    // 上一次触发时间
	triggerTimesToDay uint          // 今日已触发次数
	maxTriggerPerDate uint          // 每天最多触发次数，0 无限次数
	// maxTriggerPerHour uint          // TODO: 每小时最多触发次数，0 无限次数
}

func NewAlarm(interval time.Duration, maxTriggerPerDate uint) *Alarm {
	return &Alarm{
		interval:          interval,
		maxTriggerPerDate: maxTriggerPerDate,
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

	if a.lastTriggerAt == nil {
		shouldTrigger = true
		return
	}

	now := time.Now()

	// 如果超出了当前报警的量
	if now.Day() == a.lastTriggerAt.Day() && a.triggerTimesToDay >= a.maxTriggerPerDate {
		return false
	}

	// 如果还没有到下一次的报警时间
	if now.After(a.lastTriggerAt.Add(a.interval)) {
		shouldTrigger = true
		return
	}

	return
}
