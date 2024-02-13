package ws

import (
	"github.com/gofiber/websocket/v2"
)

func getRoom(name string) *Room {
	roomsMu.Lock()
	defer roomsMu.Unlock()

	if room, ok := rooms[name]; ok {
		return room
	}

	room := &Room{
		clients: make(map[*websocket.Conn]bool),
	}
	rooms[name] = room
	return room
}