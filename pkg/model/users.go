package model

import (
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	SocialID    string `json:"social_id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	ImageUrl    string `json:"image_url"`
	Email       string `json:"email"`
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
}

func (user *User) TableName() string {
	return "users"
}

type Credential struct {
	UserName *string `json:"user_name" valid:"Required"`
	Password *string `json:"password" valid:"Required"`
}
type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	PhoneNumber  string    `json:"phone_number"`
	DisplayName  string    `json:"display_name"`
	Id           uuid.UUID `json:"id"`
}
type CheckPhoneReq struct {
	PhoneNumber string `json:"phone_number" valid:"Required"`
}
type UserReq struct {
	ID          *uuid.UUID `json:"id" valid:"Required"`
	DisplayName *string    `json:"display_name"`
	ImageUrl    *string    `json:"image_url"`
	Password    *string    `json:"password"`
	PhoneNumber *string    `json:"phone_number"`
	Email       *string    `json:"email"`
}
type CreateUserReq struct {
	DisplayName *string `json:"display_name"`
	Password    *string `json:"password" valid:"Required"`
	PhoneNumber *string `json:"phone_number" valid:"Required"`
	Email       *string `json:"email"`
}
