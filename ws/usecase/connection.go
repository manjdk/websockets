package usecase

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/manjdk/websockets/ws"
	"golang.org/x/net/websocket"
)

type ConnectionUseCase struct {
	Conn *websocket.Conn
}

func NewConnectionUseCase(conn *websocket.Conn) ConnectionUseCase {
	return ConnectionUseCase{
		Conn: conn,
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
				log.Println("Error receiving message: ", err.Error())
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

		if err := c.sendToServer(text); err != nil {
			log.Println("Error sending message: ", err.Error())
			break
		}
	}
}

func (c *ConnectionUseCase) sendToServer(text string) error {
	message := ws.NewMessage(text)

	return websocket.JSON.Send(c.Conn, message)
}
