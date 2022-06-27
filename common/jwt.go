package common

import (
	"github.com/dgrijalva/jwt-go"
	"student-manage/model"
	"time"
)

var jwtKey = []byte("wxy-studentMS-secret_token")

type MyClaims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(manager model.Manager) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour) // 过期时间：24h
	claims := &MyClaims{
		UserId: manager.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "wxy",
			Subject:   "manager token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
