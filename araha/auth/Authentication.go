package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

var jwtKey = []byte("password")

type CustomClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func generateId() int {
	return +1
}

func GenerateToken(id, userName, role string) (string, error) {

	var user CustomClaims

	id = strconv.Itoa(generateId())
	user.Id = id

	claims := &CustomClaims{
		Id:       user.Id,
		Username: userName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtKey)
}
