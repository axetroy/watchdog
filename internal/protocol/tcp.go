package protocol

import (
	"net"
)

func PingTCP(addr string) error {
	_, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}
