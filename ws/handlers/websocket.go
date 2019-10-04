package handlers

import (
	"github.com/manjdk/websockets/ws"
	"github.com/manjdk/websockets/ws/usecase"
	"golang.org/x/net/websocket"
)

func WebSocketHandler(u usecase.WebSocketUseCase) websocket.Handler {
	return func(conn *websocket.Conn) {
		message := &ws.Message{}
		h := u.Hub

		h.AddClientChannel <- conn

		for {
			if err := websocket.JSON.Receive(conn, message); err != nil {
				message.Text = err.Error()
				h.SendChannel <- message
				u.RemoveClient(conn)
				return
			}

			h.SendChannel <- message
		}
	}
}
