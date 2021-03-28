package watchdog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	// test clear
	{
		storage := NewStorageMemory(time.Microsecond*500, time.Microsecond*500)

		storage.SetItem("foo", make([]ServiceStatus, 0))

		s := storage.GetItem("foo")

		assert.NotNil(t, s)
	}
}
