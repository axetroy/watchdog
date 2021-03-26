package protocol

import (
	"bytes"
	"context"
	"log"
	"net"
	"net/mail"
	"testing"
	"time"

	"github.com/mhale/smtpd"
	"github.com/stretchr/testify/assert"
)

func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
	return nil
}

func smtpAuthHandler(remoteAddr net.Addr, mechanism string, username []byte, password []byte, shared []byte) (bool, error) {
	isMatch := string(username) == "test" && string(password) == "test"
	return isMatch, nil
}

func rcptHandler(remoteAddr net.Addr, from string, to string) bool {
	return true
}

func CreateSMTPServer(addr string, cb func(c *smtpd.Server)) error {
	server := &smtpd.Server{
		Addr:         addr,
		Handler:      mailHandler,
		HandlerRcpt:  rcptHandler,
		AuthRequired: true,
		AuthMechs: map[string]bool{
			"LOGIN": true,
		},
		AuthHandler: smtpAuthHandler,
		Appname:     "MyServerApp",
		Hostname:    "",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	defer time.Sleep(time.Second * 2)

	time.Sleep(time.Second * 3)

	cb(server)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	return server.Shutdown(ctx)
}

func TestPingSMTP(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		args   args
		error  string
		listen bool
		port   string
		auth   interface{}
	}{
		{
			name: "2525 wrong password",
			port: "2525",
			args: args{
				addr: "localhost:2525",
			},
			error:  "535 5.7.8 Authentication credentials invalid",
			listen: true,
			auth: map[string]string{
				"username": "test",
				"password": "wrong_password",
			},
		},
		{
			name: "2424 port not listen",
			port: "2424",
			args: args{
				addr: "localhost:2424",
			},
			error:  "dial tcp [::1]:2424: connect: connection refused",
			listen: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				assert.Nil(t, CreateSMTPServer("localhost:"+tt.port, func(s *smtpd.Server) {
					err := PingSMTP(tt.args.addr, tt.auth)
					assert.EqualError(t, err, tt.error, tt.name)
				}))
			} else {
				err := PingSMTP(tt.args.addr, tt.auth)
				assert.EqualError(t, err, tt.error, tt.name)
			}

		})
	}
}

func TestPingSMTPWitoutError(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		args   args
		listen bool
		port   string
		auth   interface{}
	}{
		{
			name: "2626 right password",
			port: "2626",
			args: args{
				addr: "localhost:2626",
			},
			listen: true,
			auth: map[string]string{
				"username": "test",
				"password": "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.listen == true {
				assert.Nil(t, CreateSMTPServer("localhost:"+tt.port, func(s *smtpd.Server) {
					err := PingSMTP(tt.args.addr, tt.auth)
					assert.Nil(t, err)
				}))
			} else {
				err := PingSMTP(tt.args.addr, tt.auth)
				assert.Nil(t, err)
			}

		})
	}
}
