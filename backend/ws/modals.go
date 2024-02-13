package ws

import (
	"sync"
	"github.com/gofiber/websocket/v2"
)

type (
	Room struct {
		clients map[*websocket.Conn]bool
		mu      sync.Mutex
	}

	Rooms struct {
		ID				string	`json:"id"`
		Name 			string 	`json:"name"`
		Description		string	`json:"description"`
	}

	Message struct {
		Username string `json:"username"`
		Message  string `json:"message"`
		Type	 string `json:"type"`
	}

	Video struct {
		By 		string	`json:"by"`
		Url		string	`json:"url"`
		Type	string	`json:"type"`
	}

	Status struct {
		Username string	`json:"username"`
		Type	string	`json:"type"`
	}
)