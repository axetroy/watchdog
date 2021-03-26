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

	defer connection.Close()

	cb(connection)

	return nil
}

func TestPingTCP(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		args   args
		error  string
		listen bool
	}{
		{
			name: "9999",
			args: args{
				addr: "localhost:9999",
			},
			error:  "",
			listen: true,
		},
		{
			name: "9998",
			args: args{
				addr: "localhost:9998",
			},
			error:  "dial tcp [::1]:9998: connect: connection refused",
			listen: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				assert.Nil(t, CreateTCPServer(tt.args.addr, func(connection net.Listener) {
					err := PingTCP(context.Background(), tt.args.addr)
					if err != nil {
						assert.EqualError(t, err, tt.error, tt.name)
					} else {
						assert.Equal(t, tt.error, "", tt.name)
					}
				}))
			} else {
				err := PingTCP(context.Background(), tt.args.addr)
				if err != nil {
					assert.EqualError(t, err, tt.error, tt.name)
				} else {
					assert.Equal(t, tt.error, "", tt.name)
				}
			}

		})
	}
}
