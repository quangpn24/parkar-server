package model

import "github.com/google/uuid"

type TimeFrame struct {
	BaseModel
	Duration     int       `json:"duration"`
	Cost         float64   `json:"cost"`
	ParkingLotId uuid.UUID `json:"parking_lot_id" gorm:"not null"`
}

func (timeFrame *TimeFrame) TableName() string {
	return "time_frame"
}
