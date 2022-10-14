package jwt

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type SClaims struct {
	jwt.Claims
	Id uint64 `json:"id"`
}

// Это значение должно браться из cfg
var secret string = "secret"

func GenerateJwtById(id uint64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &SClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(100).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	tkStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("Error SignedString jwt %s !\n", err.Error())
		panic(nil)
	}

	return tkStr
}

func ParseJwt(tok string) uint64 {
	token, err := jwt.ParseWithClaims(tok, SClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Printf("Error parse jwt %s \n", err.Error())
		panic(nil)
	}

	if claims, ok := token.Claims.(*SClaims); ok && token.Valid {
		return claims.Id
	}

	fmt.Println("Error parse token !")
	return 0
}