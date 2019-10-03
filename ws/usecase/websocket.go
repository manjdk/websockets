package usecase

import (
	"strings"

	"github.com/manjdk/websockets/ws"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type WebSocketUseCase struct {
	Hub    *ws.Hub
	Logger *logrus.Logger
}

func NewWebSocketUseCase(hub *ws.Hub, logger *logrus.Logger) WebSocketUseCase {
	return WebSocketUseCase{
		Hub:    hub,
		Logger: logger,
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

	if err := send(conn, ws.NewMessage("Hello from server")); err != nil {
		w.Logger.Errorf("Failed to send welcome message: %s", err)
	}
}

func (w *WebSocketUseCase) SendMessage(message *ws.Message) {
	w.Logger.Infof("Message received: %s", message.Text)

	for _, client := range w.Hub.Client {
		message.Text = replaceQuestionToExclamation(message.Text)

		if err := send(client, message); err != nil {
			w.Logger.Errorf("Failed to send message: %s", err)
		}
	}
}

func replaceQuestionToExclamation(text string) string {
	return strings.ReplaceAll(text, "?", "!")
}
