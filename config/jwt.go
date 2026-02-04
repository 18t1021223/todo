package config

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth
var JwtAccessTokenExpiration time.Duration

func InitJWT(secret string, exp int64) {
	TokenAuth = jwtauth.New("HS256", []byte(secret), nil)
	JwtAccessTokenExpiration = time.Duration(exp)
}
