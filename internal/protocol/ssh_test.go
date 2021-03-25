package protocol

import (
	"fmt"
	"io"
	"log"
	"testing"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/stretchr/testify/assert"
	gossh "golang.org/x/crypto/ssh"
)

func CreateSSHServer(addr string, cb func(c *ssh.Server)) error {
	server := &ssh.Server{
		Addr: ":2222",
	}

	server.Handle(func(s ssh.Session) {
		authorizedKey := gossh.MarshalAuthorizedKey(s.PublicKey())
		_, _ = io.WriteString(s, fmt.Sprintf("public key used by %s:\n", s.User()))
		_, _ = s.Write(authorizedKey)
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	time.Sleep(time.Second * 3)

	defer server.Close()

	cb(server)

	return nil
}

func TestPingSSH(t *testing.T) {
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
			name: "2222",
			args: args{
				addr: "localhost:2222",
			},
			wantErr: false,
			listen:  true,
		},
		{
			name: "2223",
			args: args{
				addr: "localhost:2223",
			},
			wantErr: true,
			listen:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				err := CreateSSHServer(tt.args.addr, func(s *ssh.Server) {
					err := PingSSH(tt.args.addr, nil)
					assert.Equal(t, tt.wantErr, err != nil, tt.name)
				})

				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			} else {
				err := PingSSH(tt.args.addr, nil)
				assert.Equal(t, tt.wantErr, err != nil, tt.name)
			}

		})
	}
}
