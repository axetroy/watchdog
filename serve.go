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
	config *Config
}

type ServiceStatus struct {
	Name      string `json:"name"`
	Error     string `json:"error"`
	UpdatedAt string `json:"updated_at"`
}

func (t HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix("/api/ws", req.URL.Path) {
		s, err := socket.NewSocket(res, req)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(err.Error()))
			return
		}

		initPayload := map[string][]ServiceStatus{}

		for _, sr := range t.config.Service {
			cached := Store.GetItem(sr.Name)
			if cached == nil {
				initPayload[sr.Name] = []ServiceStatus{}
			} else {
				initPayload[sr.Name] = *cached
			}
		}

		socket.Pool.BroadcastTo(s.UUID, socket.Data{
			Event:   socket.EventInit,
			Payload: initPayload,
		})

		defer s.Close()

		for {
			_, _, err := s.ReadMessage()
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

func Serve(port string, config *Config) {
	server := http.Server{
		Addr: ":" + port,
		Handler: &HTTPHandler{
			config: config,
		},
		ReadTimeout: 3 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
