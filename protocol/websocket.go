package protocol

import (
	"golang.org/x/net/websocket"
)

func PingWebsocket(addr string) error {
	conn, err := websocket.Dial(addr, "", "http://localhost/")

	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}
