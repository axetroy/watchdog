package protocol

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateTCPServer(addr string, cb func(c net.Listener)) error {
	connection, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	cb(connection)

	return nil
}

func TestPingTCP(t *testing.T) {
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
			name: "9999",
			args: args{
				addr: "localhost:9999",
			},
			wantErr: false,
			listen:  true,
		},
		{
			name: "9998",
			args: args{
				addr: "localhost:9998",
			},
			wantErr: true,
			listen:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				err := CreateTCPServer(tt.args.addr, func(connection net.Listener) {
					defer connection.Close()
					err := PingTCP(context.Background(), tt.args.addr)
					assert.Equal(t, tt.wantErr, err != nil, tt.name)
				})

				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			} else {
				err := PingTCP(context.Background(), tt.args.addr)
				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			}

		})
	}
}
