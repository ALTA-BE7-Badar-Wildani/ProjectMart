package utilities

import (
	"time"

	entitiesDomain "go-ecommerce/entities/domain"

	"github.com/golang-jwt/jwt"
)

func CreateToken(user entitiesDomain.User) (string, error) {
	claim := jwt.MapClaims{
		"name": user.Name,
		"username": user.Username,
		"userID": user.ID,
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte("jeweteuwu"))
}