package watchdog

import (
	"embed"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/axetroy/watchdog/internal/socket"
)

//go:embed web/dist
var content embed.FS

type HTTPHandler struct {
}

func (t HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix("/api/ws", req.URL.Path) {
		socket, err := socket.NewSocket(res, req)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(err.Error()))
			return
		}

		defer socket.Close()

		for {
			_, _, err := socket.ReadMessage()
			if err != nil {
				break
			}
		}
	} else {
		req.URL.Path = "/web/dist" + req.URL.Path
		fs := http.FileServer(http.FS(content))

		fs.ServeHTTP(res, req)
	}
}

func Serve(port string) {
	server := http.Server{
		Addr:        ":" + port,
		Handler:     &HTTPHandler{},
		ReadTimeout: 3 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
