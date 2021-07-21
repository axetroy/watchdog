package watchdog

import (
	"embed"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/axetroy/watchdog/internal/socket"
)

// skipcq: SCC-compile
//go:embed web/dist
var content embed.FS

var fs = http.FileServer(http.FS(content))

type HTTPHandler struct {
	config *Config
}

type ServiceStatus struct {
	Name      string        `json:"name"`       // 服务名称
	Error     string        `json:"error"`      // 错误信息
	UpdatedAt string        `json:"updated_at"` // 检测时间
	Duration  time.Duration `json:"duration"`   // 持续时间
}

func (t HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/api/ws") {
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

		fs.ServeHTTP(res, req)
	}
}

func Serve(port string, config *Config) {
	server := http.Server{
		Addr: ":" + port,
		Handler: &HTTPHandler{
			config: config,
		},
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
