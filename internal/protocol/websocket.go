package protocol

import (
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

func PingWebsocket(addr string) error {
	conn, err := websocket.Dial(addr, "", "http://localhost/")

	if err != nil {
		return errors.WithStack(err)
	}

	defer conn.Close()

	return nil
}
