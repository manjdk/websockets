package connection

import (
	"fmt"
	"log"

	"github.com/manjdk/websockets/api/server"
	"github.com/manjdk/websockets/ws/usecase"
	"golang.org/x/net/websocket"
)

func Dial(port int) {
	container := server.NewContainer()
	conn, err := newConnection(container.HttpAddr, port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	connectionUseCase := usecase.NewConnectionUseCase(conn)
	connectionUseCase.Start()
}

func newConnection(addr string, port int) (*websocket.Conn, error) {
	return websocket.Dial(fmt.Sprintf("ws://%s:%d/ws", addr, port), "", fmt.Sprintf("http://%s", addr))
}
