package ws

import (
	"log"
	"clean/database"
)


func handleMessage(room *Room, msgData map[string]interface{}) {
	username := msgData["username"].(string)
	messageContent := msgData["message"].(string)
	messageType := msgData["type"].(string)

	user, err := database.FindUser(username)
	if err == nil {
		msg := &Message{
			Username: user.Username,
			Message:  messageContent,
			Type:     messageType,
		}
		broadcastMessage(room, msg)
	} else {
		log.Printf("User not found for ID: %s", username)
	}
}

func broadcastMessage(room *Room, msg *Message) {
	room.mu.Lock()
	defer room.mu.Unlock()

	for client := range room.clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(room.clients, client)
		}
	}
}
