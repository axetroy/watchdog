package proto

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateUDPServer(addr string, cb func(c *net.UDPConn)) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return err
	}

	connection, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		return err
	}

	cb(connection)

	return nil
}

func TestPingUDP(t *testing.T) {
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
				err := CreateUDPServer(tt.args.addr, func(connection *net.UDPConn) {
					defer connection.Close()
					err := PingUDP(tt.args.addr)
					assert.Equal(t, tt.wantErr, err != nil)
				})

				assert.Equal(t, tt.wantErr, err != nil)
			} else {
				err := PingUDP(tt.args.addr)
				assert.Equal(t, tt.wantErr, err != nil)
			}
		})
	}
}
