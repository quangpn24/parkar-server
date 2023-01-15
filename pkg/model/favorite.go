package model

import "github.com/google/uuid"

type Favorite struct {
	BaseModel
	UserId       uuid.UUID `json:"user_id"`
	ParkingLotId uuid.UUID `json:"parking_lot_id"`
}

func (f *Favorite) TableName() string {
	return "favorite"
}
