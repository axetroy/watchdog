package watchdog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name  string
		args  args
		wantR RunnerResult
	}{
		{
			name: "basic",
			args: args{
				config: Config{
					Service: []Service{
						{
							Name:     "Web",
							Protocol: "http",
							Addr:     "http://localhost:9000",
						},
						{
							Name:     "RPC",
							Protocol: "tcp",
							Addr:     "localhost:9000",
						},
					},
				},
			},
			wantR: RunnerResult{
				Errors: []RunnerError{
					{
						Name:       "Web",
						Error:      "Head \"http://localhost:9000\": dial tcp [::1]:9000: connect: connection refused",
						RetryTimes: 0,
					},
					{
						Name:       "RPC",
						Error:      "dial tcp [::1]:9000: connect: connection refused",
						RetryTimes: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Run(tt.args.config)

			assert.Equal(t, tt.wantR.Errors, result.Errors)
		})
	}
}
