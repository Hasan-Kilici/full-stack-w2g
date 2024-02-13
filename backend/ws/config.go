package ws

import (
	"sync"
	"github.com/gofiber/websocket/v2"
)

var (
	clients = make(map[*websocket.Conn]bool)
	upgrader = websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	rooms = make(map[string]*Room)
	roomsMu sync.Mutex
)
