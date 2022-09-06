package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"webapp01/utils/config"
)

type CustomClaims struct {
	string
	jwt.RegisteredClaims
}

func GenToken(userName string) (string, error) {
	fmt.Println(config.JwtConf.TokenExpired)
	claims := CustomClaims{userName, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JwtConf.TokenExpired)),
		Issuer:    config.JwtConf.Issuer,
	}}
	encryToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return encryToken.SignedString(config.JwtConf.Seed)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtConf.Seed, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
