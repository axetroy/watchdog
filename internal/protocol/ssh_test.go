package protocol

import (
	"context"
	"fmt"
	"io"
	"log"
	"testing"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/stretchr/testify/assert"
	gossh "golang.org/x/crypto/ssh"
)

func CreateSSHServer(addr string, cb func(c *ssh.Server) error) error {
	server := &ssh.Server{
		Addr: ":2222",
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			return ctx.User() == "test" && password == "test"
		},
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

	return cb(server)
}

func TestPingSSH(t *testing.T) {
	type args struct {
		addr string
		auth interface{}
	}
	tests := []struct {
		name   string
		args   args
		err    string
		listen bool
	}{
		{
			name: "2222",
			args: args{
				addr: "localhost:2222",
				auth: map[string]string{
					"username": "test",
					"password": "test",
				},
			},
			err:    "",
			listen: true,
		},
		{
			name: "2222#wrong auth",
			args: args{
				addr: "localhost:2222",
				auth: map[string]string{
					"username": "wrong_username",
					"password": "wrong_password",
				},
			},
			err:    "ssh: handshake failed: ssh: unable to authenticate, attempted methods [none password], no supported methods remain",
			listen: true,
		},
		{
			name: "2222#without password",
			args: args{
				addr: "localhost:2222",
			},
			err:    "",
			listen: true,
		},
		{
			name: "2223 if port aviable",
			args: args{
				addr: "localhost:2223",
			},
			err:    "dial tcp [::1]:2223: connect: connection refused",
			listen: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				err := CreateSSHServer(tt.args.addr, func(s *ssh.Server) error {
					return PingSSH(context.Background(), tt.args.addr, tt.args.auth)
				})

				if tt.err == "" {
					assert.Nil(t, err, tt.name)
				} else {
					assert.EqualError(t, err, tt.err, tt.name)
				}
			} else {
				err := PingSSH(context.Background(), tt.args.addr, tt.args.auth)
				if tt.err == "" {
					assert.Nil(t, err, tt.name)
				} else {
					assert.EqualError(t, err, tt.err, tt.name)
				}
			}

		})
	}
}
