package protocol

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	_, _ = io.Copy(ws, ws)
}

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: ":8888"}

	http.Handle("/echo", websocket.Handler(EchoServer))

	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

func CreateWebsocketServer(addr string, cb func()) error {
	httpServerExitDone := &sync.WaitGroup{}

	httpServerExitDone.Add(1)
	srv := startHttpServer(httpServerExitDone)

	time.Sleep(time.Second * 1)

	cb()

	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	return nil
}

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
				err := CreateWebsocketServer(tt.args.addr, func() {
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
