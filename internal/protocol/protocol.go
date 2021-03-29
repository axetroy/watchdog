package protocol

import (
	"context"

	"github.com/pkg/errors"
)

func Ping(ctx context.Context, proto string, addr string, auth interface{}) error {
	switch proto {
	case "http":
		return PingHTTP(ctx, addr)
	case "https":
		return PingHTTP(ctx, addr)
	case "tcp":
		return PingTCP(ctx, addr)
	case "udp":
		return PingUDP(addr)
	case "ws":
		return PingWebsocket(addr)
	case "wss":
		return PingWebsocket(addr)
	case "ftp":
		return PingFTP(addr)
	case "sftp":
		return PingFTP(addr)
	case "ssh":
		return PingSSH(ctx, addr, auth)
	case "smtp":
		return PingSMTP(addr, auth)
	case "pop3":
		return PingPOP3(addr)
	case "smb":
		return PingSMB(addr)
	case "nfs":
		return PingNFS(addr)
	default:
		return errors.Errorf("invalid proto '%s'", proto)
	}
}
