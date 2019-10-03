package ws

import "golang.org/x/net/websocket"

type Hub struct {
	Client              map[string]*websocket.Conn
	AddClientChannel    chan *websocket.Conn
	RemoveClientChannel chan *websocket.Conn
	SendChannel         chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Client:              make(map[string]*websocket.Conn),
		AddClientChannel:    make(chan *websocket.Conn),
		RemoveClientChannel: make(chan *websocket.Conn),
		SendChannel:         make(chan *Message),
	}
}
