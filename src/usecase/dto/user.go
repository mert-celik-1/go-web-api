package dto

import "go-web-api/src/domain/models"

type TokenDetail struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpireTime  int64
	RefreshTokenExpireTime int64
}

type RegisterUserByUsername struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func ToUserModel(from RegisterUserByUsername) models.User {
	return models.User{Username: from.Username,
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Email:     from.Email,
	}
}

type RegisterLoginByMobile struct {
	MobileNumber string
	Otp          string
}

type LoginByUsername struct {
	Username string
	Password string
}
