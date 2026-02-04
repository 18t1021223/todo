package service

import (
	"time"
	"todo_list_roadmap/config"
	db "todo_list_roadmap/db/genarated"
	"todo_list_roadmap/handle/response"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func CreateToken(u *db.User) (string, error) {
	claims := map[string]interface{}{
		// 1. "iss" (Issuer):
		"iss": "my-backend-api",

		// 2. "sub" (Subject):
		"sub": u.ID,

		// 3. "aud" (Audience):
		//"aud": "mobile-client",

		// 4. "exp" (Expiration Time)
		"exp": jwtauth.ExpireIn(time.Second * config.JwtAccessTokenExpiration),

		// 5. "nbf" (Not Before)
		//"nbf": jwtauth.EpochNow(),

		// 6. "iat" (Issued At):
		"iat": jwtauth.EpochNow(),

		// 7. "jti" (JWT ID)
		"jti": uuid.New().String(),

		"u_email": u.Email,
	}

	_, s, err := config.TokenAuth.Encode(claims)
	if err != nil {
		zap.L().Info("Failed CreateToken", zap.Error(err), zap.String("ID", u.ID))
		return "", response.ErrInternalServer
	}
	return s, nil
}
