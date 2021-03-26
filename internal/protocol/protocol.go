package protocol

import (
	"context"

	"github.com/pkg/errors"
)

func Ping(proto string, addr string, auth interface{}, ctx context.Context) error {
	switch proto {
	case "http":
		fallthrough
	case "https":
		return PingHTTP(ctx, addr)
	case "tcp":
		return PingTCP(ctx, addr)
	case "udp":
		return PingUDP(addr)
	case "ws":
		fallthrough
	case "wss":
		return PingWebsocket(addr)
	case "ftp":
		fallthrough
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
