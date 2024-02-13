package middleware

import (
	"clean/utils"
	jwtware "github.com/gofiber/contrib/jwt"
)

var Jwt = jwtware.New(jwtware.Config{
	KeyFunc: utils.CustomKeyFunc(),
})