package database

import "log"

func InsertRoom(room Room) {
	query := "INSERT INTO rooms (id, name, description, public) VALUES (?, ?, ?, ?);"
	_, err := db.Exec(query, room.ID, room.Name, room.Description, room.Public)
	if err != nil {
		log.Fatal(err)
	}
}

func ListRooms() []Room {
	var rooms []Room
	rows, err := db.Query("SELECT id, name, description FROM rooms WHERE public = true;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var room Room
		err := rows.Scan(&room.ID, &room.Name, &room.Description)
		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms
}

func FindRoom(roomID int64) (Rooms, error) {
    var room Rooms

    query := `
        SELECT r.id, r.name, r.description, rs.stopVideo, rs.changeVideo, rs.videoRequest
        FROM rooms r
        LEFT JOIN roomsettings rs ON r.id = rs.roomID
        WHERE r.id = ?;
    `
    err := db.QueryRow(query, roomID).Scan(&room.ID, &room.Name, &room.Description, &room.StopVideo, &room.ChangeVideo, &room.VideoRequest)
    if err != nil {
        return room, err
    }

    return room, nil
}


func DeleteRoom(roomID int64) error {
	query := "DELETE FROM rooms WHERE id = ?;"
	_, err := db.Exec(query, roomID)
	if err != nil {
		return err
	}
	return nil
}

func InsertRoomMember(roomMember RoomMember) {
	query := "INSERT INTO roommembers (roomID, userID, username, perm) VALUES (?, ?, ?, ?);"
	_, err := db.Exec(query, roomMember.RoomID, roomMember.UserID, roomMember.Username, roomMember.Perm)
	if err != nil {
		log.Fatal(err)
	}
}

func FindRoomMember(roomID int64, userID string) (RoomMember, error) {
	var roomMember RoomMember
	query := "SELECT roomID, userID, username, perm FROM roommembers WHERE roomID = ? AND userID = ?;"
	err := db.QueryRow(query, roomID, userID).Scan(&roomMember.RoomID, &roomMember.UserID, &roomMember.Username, &roomMember.Perm)
	if err != nil {
		return roomMember, err
	}
	return roomMember, nil
}

func ListRoomMembers(roomID int64) []RoomMember {
	var roomMembers []RoomMember
	query := "SELECT roomID, userID, username, perm FROM roommembers WHERE roomID = ?;"
	rows, err := db.Query(query, roomID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var roomMember RoomMember
		err := rows.Scan(&roomMember.RoomID, &roomMember.UserID, &roomMember.Username, &roomMember.Perm)
		if err != nil {
			log.Fatal(err)
		}
		roomMembers = append(roomMembers, roomMember)
	}
	return roomMembers
}

func DeleteRoomMember(roomID int64, userID string) error {
	query := "DELETE FROM roomMembers WHERE roomID = ? AND userID = ?;"
	_, err := db.Exec(query, roomID, userID)
	if err != nil {
		return err
	}
	return nil
}

func InsertRoomSetting(roomSetting RoomSetting) {
	query := "INSERT INTO roomsettings (roomID, stopVideo, changeVideo, videoRequest) VALUES (?, ?, ?, ?);"
	_, err := db.Exec(query, roomSetting.RoomID, roomSetting.StopVideo, roomSetting.ChangeVideo, roomSetting.VideoRequest)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteRoomSetting(roomID int64) error {
	query := "DELETE FROM roomsettings WHERE roomID = ?;"
	_, err := db.Exec(query, roomID)
	if err != nil {
		return err
	}
	return nil
}