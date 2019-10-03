package server

import (
	"github.com/manjdk/websockets/config"
	"github.com/manjdk/websockets/ws"
	"github.com/manjdk/websockets/ws/usecase"
)

type Container struct {
	Hub              *ws.Hub
	WebsocketUseCase usecase.WebSocketUseCase
	HttpAddr         string
}

func NewContainer() *Container {
	config.InitViper()
	hub := ws.NewHub()

	webSocketUseCase := usecase.NewWebSocketUseCase(hub)

	return &Container{
		Hub:              hub,
		WebsocketUseCase: webSocketUseCase,
		HttpAddr:         config.HttpAddr(),
	}
}
