package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/golang-jwt/jwt"
)

func CreateJWT(id string) (string, error) {

	mySigningKey := []byte("Secret") // token 키

	token := jwt.New(jwt.SigningMethodHS256) // token 프로토콜

	claims := token.Claims.(jwt.MapClaims) // token정보  id, 유효기간
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	tk, err := token.SignedString(mySigningKey) // token키를 가지고 sign

	if err != nil {
		return "", err
	}

	return tk, nil
}
