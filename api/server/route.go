package server

import (
	"net/http"

	"github.com/manjdk/websockets/ws/handlers"
)

func Route(router *http.ServeMux, c *Container) *http.ServeMux {
	router.Handle("/ws", handlers.WebSocketHandler(c.WebsocketUseCase))

	return router
}
