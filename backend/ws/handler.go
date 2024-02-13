package ws

import (
	"log"
	"github.com/gofiber/websocket/v2"
)


func HandleRooms(c *websocket.Conn, roomName string) {
	room := getRoom(roomName)

	room.mu.Lock()
	room.clients[c] = true
	room.mu.Unlock()

	defer func() {
		room.mu.Lock()
		delete(room.clients, c)
		room.mu.Unlock()
		c.Close()
	}()

	for {
		var msgData map[string]interface{}
		err := c.ReadJSON(&msgData)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		switch msgType := msgData["type"].(string); msgType {
		case "message":
			handleMessage(room, msgData)
		case "change_video", "request_video":
			handleVideo(room, msgData)
		case "join", "leave":
			handleStatus(room, msgData)
		default:
			log.Printf("Unknown message type: %s", msgType)
		}
	}
}