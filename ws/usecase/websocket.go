package usecase

import (
	"log"
	"strings"

	"github.com/manjdk/websockets/ws"
	"golang.org/x/net/websocket"
)

type WebSocketUseCase struct {
	Hub *ws.Hub
}

func NewWebSocketUseCase(hub *ws.Hub) WebSocketUseCase {
	return WebSocketUseCase{
		Hub: hub,
	}
}

func (w *WebSocketUseCase) Start() {
	h := w.Hub

	for {
		select {
		case conn := <-h.AddClientChannel:
			w.AddClient(conn)
		case conn := <-h.RemoveClientChannel:
			w.RemoveClient(conn)
		case m := <-h.SendChannel:
			w.SendMessage(m)
		}
	}
}

func (w *WebSocketUseCase) RemoveClient(conn *websocket.Conn) {
	delete(w.Hub.Client, conn.LocalAddr().String())
}

func (w *WebSocketUseCase) AddClient(conn *websocket.Conn) {
	w.Hub.Client[conn.RemoteAddr().String()] = conn
	w.send(conn, ws.NewMessage("Hello from server"))
}

func (w *WebSocketUseCase) SendMessage(message *ws.Message) {
	log.Printf("Message received: %s", message.Text)

	for _, client := range w.Hub.Client {
		message.Text = strings.ReplaceAll(message.Text, "?", "!")

		w.send(client, message)
	}
}

func (w *WebSocketUseCase) send(conn *websocket.Conn, message *ws.Message) {
	if err := websocket.JSON.Send(conn, message); err != nil {
		log.Println("Error broadcasting message: ", err)
		return
	}
}
