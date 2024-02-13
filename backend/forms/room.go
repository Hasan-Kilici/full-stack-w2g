package forms

type CreateRoom struct {
	Name			string	`json:"name" validate:"min=5,max=20,required"`
	Description		string	`json:"description" validate:"max=200,required"`
	IsPublic		bool	`json:"public"`
	StopVideo		bool	`json:"stopVideo"`
	ChangeVideo		bool	`json:"changeVideo"`
	VideoRequest	bool	`json:"videoRequest"`
}

type JoinRoom struct {
	RoomID		    string	`json:"roomID"`
	UserID			string	`json:"userID"`
	Username		string	`json:"username"`
	Perm			string	`json:"perm"`
}

type LeaveRoom struct {
	RoomID		    string	`json:"roomID"`
	UserID			string	`json:"userID"`
}