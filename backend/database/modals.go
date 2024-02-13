package database

type (
	Rooms struct {
		Room
		RoomSetting
	}

	Room struct {
		ID          int64
		Name        string
		Description string
		Public		bool
	}

	RoomSetting struct {
		RoomID       int64
		StopVideo    bool
		ChangeVideo  bool
		VideoRequest bool
	}

	User struct {
		Token	 string
		ID       string		`json:"id"`
		Username string		`json:"username"`
	}

	RoomMember struct {
		RoomID 		int64
		UserID 		string
		Username	string
		Perm   		string
	}
)