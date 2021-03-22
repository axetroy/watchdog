package socket

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Data struct {
	Event   Event       `json:"event"`
	Payload interface{} `json:"payload"`
}

type Socket struct {
	UUID string
	conn *websocket.Conn
}

func NewSocket(res http.ResponseWriter, req *http.Request) (*Socket, error) {
	conn, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	uid := uuid.New()

	socket := &Socket{
		conn: conn,
		UUID: uid.String(),
	}

	Pool.Add(uid.String(), socket)

	return socket, nil
}

func (s *Socket) ReadMessage() (messageType int, p []byte, err error) {
	return s.conn.ReadMessage()
}

func (s *Socket) WriteMessage(messageType int, data []byte) error {
	return s.conn.WriteMessage(messageType, data)
}

func (s *Socket) WriteJSON(v Data) error {
	return s.conn.WriteJSON(v)
}

func (s *Socket) Close() error {
	defer func() {
		Pool.Remove(s.UUID)
	}()

	if err := s.conn.Close(); err != nil {
		return err
	}

	return nil
}
