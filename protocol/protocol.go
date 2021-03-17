package protocol

import (
	"context"

	"github.com/pkg/errors"
)

func Ping(proto string, addr string, ctx context.Context) error {
	switch proto {
	case "http":
		fallthrough
	case "https":
		return PingHTTP(addr)
	case "tcp":
		return PingTCP(addr)
	case "udp":
		return PingUDP(addr)
	case "ws":
		fallthrough
	case "wss":
		return PingWebsocket(addr)
	default:
		return errors.Errorf("invalid proto '%s'", proto)
	}
}
