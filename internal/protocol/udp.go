package protocol

import (
	"net"

	"github.com/pkg/errors"
)

func PingUDP(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return errors.WithStack(err)
	}

	// 连接服务器
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		return errors.WithStack(err)
	}

	defer conn.Close()

	return nil
}
