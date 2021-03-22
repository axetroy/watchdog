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
		alarm := NewAlarm(time.Millisecond*200, 20) // 间隔 2s，每天最多触发次数 20 次

		index := 0
		for {
			if index == 20 {
				break
			}

			index++

			_ = alarm.ShouldTrigger()

			// The default quantum under Linux is 10ms so this is expected behavior and is a property of Linux, not go.
			time.Sleep(time.Millisecond * 105)
		}

		assert.Equal(t, 10, int(alarm.triggerTimesToDay), index)
	}
}
