package validate

import (
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pradeep/golang-micro/config"
	"github.com/pradeep/golang-micro/model"
)

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	pattern := regexp.MustCompile(regex)
	return pattern.MatchString(email)
}

func ValidateToken(signedtoken string) (claims *model.SignedTokens, err error) {
	token, err := jwt.ParseWithClaims(signedtoken, &model.SignedTokens{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env.Jwt_SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.SignedTokens)
	if !ok {
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}
	return claims, nil
}
