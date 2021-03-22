package protocol

import (
	"net"
)

func PingUDP(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return err
	}

	// 连接服务器
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}
