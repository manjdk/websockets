package usecase

import (
	"github.com/manjdk/websockets/ws"
	"golang.org/x/net/websocket"
)

func send(conn *websocket.Conn, message *ws.Message) error {
	if err := websocket.JSON.Send(conn, message); err != nil {
		return err
	}

	return nil
}
