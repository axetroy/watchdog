package notify

import (
	"sort"
	"strings"
	"testing"

	"github.com/axetroy/watchdog"
	"github.com/axetroy/watchdog/internal/tester"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	type args struct {
		content  string
		reporter watchdog.Reporter
	}
	tests := []struct {
		name  string
		args  args
		error string
	}{
		{
			name: "local server",
			args: args{
				content: "test",
				reporter: watchdog.Reporter{
					Protocol: "webhook",
					Target:   []string{"http://localhost:3030"},
				},
			},
			error: "",
		},
		{
			name: "unknown server",
			args: args{
				content: "test",
				reporter: watchdog.Reporter{
					Protocol: "webhook",
					Target:   []string{"http://localhost:12345"},
				},
			},
			error: "Post \"http://localhost:12345\": dial tcp [::1]:12345: connect: connection refused",
		},
		{
			name: "multiple target server",
			args: args{
				content: "test",
				reporter: watchdog.Reporter{
					Protocol: "webhook",
					Target:   []string{"http://localhost:3030", "http://localhost:12345"},
				},
			},
			error: "Post \"http://localhost:12345\": dial tcp [::1]:12345: connect: connection refused",
		},
		{
			name: "multiple target server and all error",
			args: args{
				content: "test",
				reporter: watchdog.Reporter{
					Protocol: "webhook",
					Target:   []string{"http://localhost:54321", "http://localhost:12345"},
				},
			},
			error: "Post \"http://localhost:12345\": dial tcp [::1]:12345: connect: connection refused;Post \"http://localhost:54321\": dial tcp [::1]:54321: connect: connection refused",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, tester.CreateHTTPEchoServer(":3030", func() {
				err := Webhook(tt.args.content, tt.args.reporter)
				if tt.error != "" {
					arr := strings.Split(err.Error(), ";")
					sort.Strings(arr)
					assert.Equal(t, strings.Join(arr, ";"), tt.error)
				} else {
					assert.Nil(t, err)
				}
			}))
		})
	}
}
