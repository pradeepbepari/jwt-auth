package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserRegister struct {
	FirstName string `jsom:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Uuid       uuid.UUID
	FirstName  string
	LastName   string
	Password   string
	Email      string
	Phone      string
	Role       string
	User_id    string
	Created_at time.Time
	Updated_at time.Time
}
type Users struct {
	User_id    string
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	Role       string
	Created_at time.Time
	Updated_at time.Time
}
type SignedTokens struct {
	Email     string
	FirstName string
	Id        string
	Role      string
	jwt.StandardClaims
}
