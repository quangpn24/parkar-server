package model

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
