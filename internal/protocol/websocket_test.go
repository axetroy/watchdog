package protocol

import (
	"testing"

	"github.com/axetroy/watchdog/internal/tester"
	"github.com/stretchr/testify/assert"
)

func TestPingWebsocket(t *testing.T) {
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
				addr: "ws://localhost:8888/echo",
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
				err := tester.CreateWebsocketServer(tt.args.addr, func() {
					err := PingWebsocket(tt.args.addr)
					assert.Equal(t, tt.wantErr, err != nil, tt.name)
				})

				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			} else {
				err := PingWebsocket(tt.args.addr)
				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			}
		})
	}
}
