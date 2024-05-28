package auth

import (
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/golang-jwt/jwt/v4"
)

type RegisterUserReq struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Email    string `json:"email"    validate:"required,email"`
	Role     role.Role
}

type RegisterUserRes struct {
	AccessToken string `json:"token"`
}

type LoginUserReq struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Role     role.Role
}

type LoginUserRes struct {
	AccessToken string `json:"token"`
}

type GetSessionFromTokenReq struct {
	AccessToken string
}

type GetSessionFromTokenRes struct {
	UserId string
	Role   role.Role
}

type jwtSubClaims struct {
	UserId string    `json:"userId"`
	Role   role.Role `json:"role"`
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Data jwtSubClaims `json:"data"`
}
