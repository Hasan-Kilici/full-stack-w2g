package ws

import (
	"log"
	"clean/database"
)

func handleStatus(room *Room, msgData map[string]interface{}) {
	username := msgData["username"].(string)
	statusType := msgData["type"].(string)

	user, err := database.FindUser(username)
	if err == nil {
		msg := &Status{
			Username: user.Username,
			Type:     statusType,
		}
		broadcastStatus(room, msg)
	} else {
		log.Printf("User not found for Token: %s", username)
	}
}

func broadcastStatus(room *Room, msg *Status) {
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
