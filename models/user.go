package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
