package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type JwtTokenStruct struct {
	jwtSecret []byte
}

func (j JwtTokenStruct) Parse(tokenString string) (*jwt.Token, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return j.jwtSecret, nil
	})

	if err != nil {
		log.Fatalln(err.Error())
		return token, err
	}

	return token, err
}

func (j JwtTokenStruct) GetUser(token *jwt.Token) (*uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, http.ErrAbortHandler
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, http.ErrAbortHandler
		}
	} else {
		return nil, http.ErrAbortHandler
	}

	userID := uint(claims["user_id"].(float64))
	return &userID, nil
}

func (j JwtTokenStruct) Create(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expires in 30 day
	})

	return token.SignedString(j.jwtSecret)
}

var JwtToken = JwtTokenStruct{jwtSecret: []byte(os.Getenv("JWT_SECRET"))}
