package protocol

import (
	"context"

	"github.com/pkg/errors"
)

// 检测服务 TODO: 使用 ctx 进行超时检测
func Ping(proto string, addr string, auth interface{}, ctx context.Context) error {
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
	case "ftp":
		fallthrough
	case "sftp":
		return PingFTP(addr)
	case "ssh":
		return PingSSH(addr, auth)
	case "smtp":
		return PingSMTP(addr)
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
