package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pradeep/golang-micro/config"
	"github.com/pradeep/golang-micro/model"
)

func GenerateAllToken(userid, name, email, role string) (token, refreshtoken string, err error) {
	var secretkey = []byte(config.Env.Jwt_SecretKey)
	claims := &model.SignedTokens{
		FirstName: name,
		Id:        userid,
		Email:     email,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	refreshClaims := &model.SignedTokens{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}
	accesstoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretkey)
	if err != nil {
		return "", "", fmt.Errorf("error:Accessing token")
	}
	refresh_tokens, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretkey)
	if err != nil {
		return "", "", fmt.Errorf("error:Accessing Refreshtoken")
	}

	return accesstoken, refresh_tokens, nil
}
