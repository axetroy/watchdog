package protocol

import (
	"testing"

	"github.com/axetroy/watchdog/internal/tester"
	"github.com/stretchr/testify/assert"
)

func TestPingFTP(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		listen  bool
	}{
		{
			name: "8888",
			args: args{
				addr: "localhost:8888",
			},
			wantErr: false,
			listen:  true,
		},
		{
			name: "8887",
			args: args{
				addr: "localhost:8887",
			},
			wantErr: true,
			listen:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				err := tester.CreateFTPServer(tt.args.addr, func() {
					err := PingFTP(tt.args.addr)
					assert.Equal(t, tt.wantErr, err != nil, tt.name)
				})
				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			} else {
				err := PingFTP(tt.args.addr)
				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			}

		})
	}
}
