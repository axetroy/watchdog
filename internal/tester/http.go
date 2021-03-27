package tester

import (
	"context"
	"io"
	"net/http"
	"time"
)

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, r.Body)
	})
}

func CreateHTTPEchoServer(addr string, cb func()) error {
	server := &http.Server{
		Addr:    addr,
		Handler: handler(),
	}

	go func() {
		_ = server.ListenAndServe()
	}()

	time.Sleep(time.Second * 3)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	defer func() {
		_ = server.Shutdown(ctx)
	}()

	cb()

	return nil
}
