package main

import (
	"github.com/dgrijalva/jwt-go"
)

var SingedToken string

func CreateToken(claims *Bang) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func CheckAction(strToken string) (*Bang, error) {
	token, err := jwt.ParseWithClaims(strToken, &Bang{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Bang)
	if !ok {
		return nil, err
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	//fmt.Println("verify")
	return claims, nil
}