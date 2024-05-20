package authenticate

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pradeep/golang-micro/config"
	"github.com/pradeep/golang-micro/model"
	"github.com/pradeep/golang-micro/utils"
)

func Authenticated(req http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("jwt-token")
		claims := model.SignedTokens{}
		tokens, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Env.Jwt_SecretKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token signature"))
				return
			}
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		if !tokens.Valid {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("token has expired"))
			return
		}

		r.Header.Set("email", claims.Email)
		r.Header.Set("firstname", claims.FirstName)
		r.Header.Set("id", claims.Id)
		r.Header.Set("role", claims.Role)

		req(w, r)
	}
}
