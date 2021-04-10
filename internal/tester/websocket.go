package tester

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

func websocketEchoHandler(ws *websocket.Conn) {
	_, _ = io.Copy(ws, ws)
}

func startWebsocketServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: ":8888"}

	http.Handle("/echo", websocket.Handler(websocketEchoHandler))

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
	srv := startWebsocketServer(httpServerExitDone)

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
