package notify

import (
	"testing"

	"github.com/axetroy/watchdog"
	"github.com/stretchr/testify/assert"
)

func TestWechat(t *testing.T) {
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
			name: "token error",
			args: args{
				content: "test report",
				reporter: watchdog.Reporter{
					Protocol: "wechat",
					Target:   []string{"test_uid"},
					Payload: map[string]string{
						"app_token": "123123",
					},
				},
			},
			error: "1001 业务异常错误 appToken不正确",
		},
		{
			name: "token valid",
			args: args{
				content: "test report",
				reporter: watchdog.Reporter{
					Protocol: "wechat",
					Target:   []string{"test_uid"},
					Payload: map[string]string{
						"app_token": "AT_RveipVXCuosGYFH7Q1MLC5wkdtfyxFjM", // this is a real app token
					},
				},
			},
			error: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wechat(tt.args.content, tt.args.reporter)
			if tt.error != "" {
				assert.EqualError(t, err, tt.error)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
