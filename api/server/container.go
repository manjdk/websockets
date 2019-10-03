package server

import (
	"github.com/manjdk/websockets/config"
	"github.com/manjdk/websockets/ws"
	"github.com/manjdk/websockets/ws/usecase"
	"github.com/sirupsen/logrus"
)

type Container struct {
	Hub              *ws.Hub
	WebsocketUseCase usecase.WebSocketUseCase
	HttpAddr         string
	Logger           *logrus.Logger
}

func NewContainer() *Container {
	config.InitViper()
	logger := logrus.New()
	hub := ws.NewHub()

	webSocketUseCase := usecase.NewWebSocketUseCase(hub, logger)

	return &Container{
		Hub:              hub,
		WebsocketUseCase: webSocketUseCase,
		HttpAddr:         config.HttpAddr(),
		Logger:           logger,
	}
}
