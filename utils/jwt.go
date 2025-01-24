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

var AccessTokenKey = []byte("acm=ave_mujica+crychic+mygo")
var RefreshTokenKey = []byte("oiiaiio...")

func SetAccessToken(uid int) (string, error) {
	AccessClaim := MyClaims{
		Uid: uid,
		MapClaims: jwt.MapClaims{
			"iss": "ecommerce",
			"exp": jwt.NewNumericDate(time.Now().Add(time.Hour)),
			"nbf": jwt.NewNumericDate(time.Now().Add(time.Second * -1)),
			"iat": jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessClaim)
	return accessToken.SignedString(AccessTokenKey)
}

func SetRefreshToken(uid int) (string, error) {
	RefreshClaim := MyClaims{
		Uid: uid,
		MapClaims: jwt.MapClaims{
			"iss": "ecommerce",
			"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			"nbf": jwt.NewNumericDate(time.Now()),
			"iat": jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaim)
	return refreshToken.SignedString(RefreshTokenKey)
}
func ParseAccessToken(t string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return AccessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid accessToken")
}

func ParseRefreshToken(t string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid refreshToken")
}
