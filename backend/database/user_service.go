package database

import (
	"log"
	"github.com/patrickmn/go-cache"
)

func InsertUser(user User) {
	query := "INSERT INTO users (token, id, username) VALUES (?, ?, ?);"
	_, err := db.Exec(query, user.Token, user.ID, user.Username)
	if err != nil {
		log.Fatal(err)
	}
}

func FindUser(token string) (User, error) {
	if user, found := userCache.Get(token); found {
        return user.(User), nil
    }

    var user User
    query := "SELECT id, username FROM users WHERE token = ?;"
    err := db.QueryRow(query, token).Scan(&user.ID, &user.Username)
    if err != nil {
        return user, err
    }

    userCache.Set(token, user, cache.DefaultExpiration)

    return user, nil
}

func FindUserById(id string) (User, error) {
	if user, found := userCache.Get(id); found {
        return user.(User), nil
    }

    var user User
    query := "SELECT token, id, username FROM users WHERE id = ?;"
    err := db.QueryRow(query, id).Scan(&user.Token, &user.ID, &user.Username)
    if err != nil {
        return user, err
    }

    userCache.Set(id, user, cache.DefaultExpiration)

    return user, nil
}

func DeleteUser(userID string) error {
	query := "DELETE FROM users WHERE id = ?;"
	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}