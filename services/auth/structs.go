package auth

import "github.com/golang-jwt/jwt/v4"

type RegisterUserReq struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Email    string `json:"email"    validate:"required,email"`
	Role     string `                validate:"required,oneof=admin user" param:"role"`
}

type RegisterUserRes struct {
	AccessToken string `json:"token"`
}

type LoginUserReq struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Role     string `                validate:"required,oneof=admin user" param:"role"`
}

type LoginUserRes struct {
	AccessToken string `json:"token"`
}

type GetSessionFromTokenReq struct {
	AccessToken string
}

type GetSessionFromTokenRes struct {
	UserId  string
	IsAdmin bool
}

type jwtSubClaims struct {
	UserId  string `json:"userId"`
	IsAdmin bool   `json:"is_admin"`
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Data jwtSubClaims `json:"data"`
}
