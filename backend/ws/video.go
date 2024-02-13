package ws

import (
	"log"
	"clean/database"
)


func sendVideoMessage(room *Room, msg *Video) {
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

func handleChangeVideo(room *Room, msg *Video) {
	sendVideoMessage(room, msg)
}

func handleVideo(room *Room, msgData map[string]interface{}) {
	by := msgData["by"].(string)
	url := msgData["url"].(string)
	videoType := msgData["type"].(string)

	user, err := database.FindUser(by)
	if err == nil {
		msg := &Video{
			By:   user.Username,
			Url:  url,
			Type: videoType,
		}
		sendVideoMessage(room, msg)
	} else {
		log.Printf("User not found for Token: %s", by)
	}
}

func broadcastVideo(room *Room, msg *Video) {
	sendVideoMessage(room, msg)
}
