package utils

import (
	"TikTok_Project/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySignKey = []byte("bjxf.douyin")

type Claims struct{
	UserId int64
	jwt.StandardClaims
}
// 获取 token
func GenToken(user repository.User) (string, error){
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "douyin_pro_bjxf",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(mySignKey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
// 解析 token
func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySignKey, nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}
	return nil, false
}