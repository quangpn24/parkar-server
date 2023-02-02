package model

import "github.com/google/uuid"

type Favorite struct {
	BaseModel
	UserId       uuid.UUID `json:"user_id"`
	ParkingLotId uuid.UUID `json:"parking_lot_id"`
	//ParkingLot ParkingLot
}

func (f *Favorite) TableName() string {
	return "favorite"
}

type FavoriteRequest struct {
	UserId       uuid.UUID `json:"user_id" valid:"Required"`
	ParkingLotId uuid.UUID `json:"parking_lot_id" valid:"Required"`
}
