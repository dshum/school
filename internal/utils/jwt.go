package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dshum/school/internal/config"
	"time"
)

func CreateToken(userId int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(config.JWT.TTL)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(jwtToken string) (*jwt.Token, error) {
	return jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWT.Secret), nil
	})
}

func CreateAuth(id int, token string) error {
	return errors.New("error")
}
