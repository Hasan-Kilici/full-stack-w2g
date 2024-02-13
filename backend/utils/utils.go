package utils

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"time"
)

func GenerateToken(Snowflake string) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, 12)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}

	encoding := base32.StdEncoding

	snowflake := encoding.EncodeToString([]byte(Snowflake))
	timestamp := GenerateID()

	token := fmt.Sprintf("%s.%s.%s", snowflake, timestamp, string(salt))

	return token
}

func GenerateID() string {
	currentTime := time.Now()

	year := currentTime.Year()
	month := int(currentTime.Month())
	day := currentTime.Day()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()

	data := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	encoded := base32.StdEncoding.EncodeToString([]byte(data))

	return encoded
}
