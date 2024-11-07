package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	Id int `json:"id"`
}

var SECRET_KEY, _ = os.LookupEnv("SECRET_KEY")

func CreateToken(id uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (int, error) {
	ptr := new(AuthClaims)
	token, err := jwt.ParseWithClaims(tokenString, ptr, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		return 0, err
	} else if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	} else {
		// claims := token.Claims.(*AuthClaims)
		claims := *ptr
		return claims.Id, nil
	}
}
