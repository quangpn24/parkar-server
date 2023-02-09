package model

import "github.com/google/uuid"

type Favorite struct {
	BaseModel
	UserId       uuid.UUID `json:"userId" gorm:"type:uuid"`
	ParkingLotId uuid.UUID `json:"parkingLotId" gorm:"type:uuid"`
	//ParkingLot ParkingLot
}

func (f *Favorite) TableName() string {
	return "favorite"
}

type FavoriteRequest struct {
	UserId       uuid.UUID `json:"userId" valid:"Required"`
	ParkingLotId uuid.UUID `json:"parkingLotId" valid:"Required"`
}
