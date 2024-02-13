package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	jwtware "github.com/gofiber/contrib/jwt"
)

func CustomKeyFunc() jwt.Keyfunc {
    return func(t *jwt.Token) (interface{}, error) {
        if t.Method.Alg() != jwtware.HS256 {
            return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
        }

   		signingKey := "secret"

        return []byte(signingKey), nil
    }
}