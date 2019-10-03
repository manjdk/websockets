package usecase

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manjdk/websockets/ws"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type ConnectionUseCase struct {
	Conn   *websocket.Conn
	Logger *logrus.Logger
}

func NewConnectionUseCase(conn *websocket.Conn, logger *logrus.Logger) ConnectionUseCase {
	return ConnectionUseCase{
		Conn:   conn,
		Logger: logger,
	}
}

func (c *ConnectionUseCase) Start() {
	c.receiveMessages()
	c.scanAndSendMessage()
}

func (c *ConnectionUseCase) receiveMessages() {
	message := &ws.Message{}

	go func() {
		for {
			if err := websocket.JSON.Receive(c.Conn, message); err != nil {
				c.Logger.Errorf("Error receiving message: %s", err)
				break
			}

			fmt.Println("Response: ", message.Text)
		}
	}()
}

func (c *ConnectionUseCase) scanAndSendMessage() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		message := ws.NewMessage(text)

		if err := send(c.Conn, message); err != nil {
			c.Logger.Errorf("Error sending message: %s", err)
			break
		}
	}
}
