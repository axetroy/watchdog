package watchdog

import (
	"log"
	"net/http"
	"time"
)

type HTTPHandler struct {
}

func (t HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("foo", "bar")
	_, _ = res.Write([]byte("hello world"))
}

func Serve(port string) {
	server := http.Server{
		Addr:        ":" + port,
		Handler:     &HTTPHandler{},
		ReadTimeout: 3 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
