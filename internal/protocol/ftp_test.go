package protocol

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"goftp.io/server/v2"
)

func CreateFTPServer(addr string, cb func(c *server.Server)) error {
	server, err := server.NewServer(&server.Options{
		Port: 8888,
		Perm: server.NewSimplePerm("test", "test"),
	})

	if err != nil {
		return err
	}

	defer func() {
		_ = server.Shutdown()
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	time.Sleep(time.Second * 3)

	cb(server)

	return nil
}

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
				err := CreateFTPServer(tt.args.addr, func(connection *server.Server) {
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
