package protocol

import (
	"github.com/pkg/errors"
)

func Ping(proto string, addr string) error {
	switch proto {
	case "http":
	case "https":
		return PingHTTP(addr)
	case "tcp":
		return PingTCP(addr)
	case "udp":
		return PingUDP(addr)
	case "ws":
	case "wss":
		return PingWebsocket(addr)
	}

	return errors.Errorf("invalid proto '%s'", proto)
}
