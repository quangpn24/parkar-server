package model

import "github.com/google/uuid"

type Company struct {
	BaseModel
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
}

func (company *Company) TableName() string {
	return "company"
}

type CompanyReq struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"companyName" valid:"Required"`
	PhoneNumber *string    `json:"phoneNumber" valid:"Required"`
	Email       *string    `json:"email" valid:"Required"`
	Password    *string    `json:"password" valid:"Required"`
}

type LoginReq struct {
	Email    *string `json:"email" valid:"Required"`
	Password *string `json:"password" valid:"Required"`
}
