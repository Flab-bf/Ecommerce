package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	Uid int `json:"uid"`
	jwt.MapClaims
}

var MySecretKey = []byte("acm=ave_mujica+crychic+mygo")

func SetTokenJwt(uid int, duration time.Duration) (string, error) {
	Claim := MyClaims{
		Uid: uid,
		MapClaims: jwt.MapClaims{
			"iss": "ecommerce",
			"exp": jwt.NewNumericDate(time.Now().Add(duration)),
			"nbf": jwt.NewNumericDate(time.Now().Add(time.Second * -1)),
			"iat": jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	return token.SignedString(MySecretKey)
}

func ParseToken(t string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func RefreshToken(uid int) (string, error) {
	return SetTokenJwt(uid, time.Hour*24)
}
