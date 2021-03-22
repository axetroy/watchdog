package watchdog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlarm(t *testing.T) {
	{
		alarm := NewAlarm(time.Microsecond*500, 20) // 间隔 0.5s，每天最多触发次数 20 次

		index := 0
		for {
			if index == 20 {
				break
			}

			assert.True(t, alarm.ShouldTrigger(), index)
			assert.Equal(t, int(uint(index+1)), int(alarm.triggerTimesToDay), index)

			index++

			time.Sleep(time.Microsecond * 600)
		}

		assert.False(t, alarm.ShouldTrigger(), index)
		assert.Equal(t, 20, int(alarm.triggerTimesToDay), index)

		assert.False(t, alarm.ShouldTrigger(), index)
		assert.Equal(t, 20, int(alarm.triggerTimesToDay), index)

		assert.False(t, alarm.ShouldTrigger(), index)
		assert.Equal(t, 20, int(alarm.triggerTimesToDay), index)
	}

	{
		alarm := NewAlarm(time.Microsecond*200, 20) // 间隔 2s，每天最多触发次数 20 次

		index := 0
		for {
			if index == 20 {
				break
			}

			index++

			_ = alarm.ShouldTrigger()

			time.Sleep(time.Microsecond * 100)
		}

		assert.Equal(t, 11, int(alarm.triggerTimesToDay), index)
	}
}
