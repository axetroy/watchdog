package protocol

import (
	"net"
)

func PingUDP(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		return nil
	}

	defer conn.Close()

	return nil
}
