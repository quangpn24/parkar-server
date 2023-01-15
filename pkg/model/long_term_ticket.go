package model

import (
	"github.com/google/uuid"
	"time"
)

type LongTermTicket struct {
	BaseModel
	Type          string     `json:"type"`
	StartTime     *time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
	VehicleId     uuid.UUID  `json:"vehicle_id"`
	ParkingLotId  uuid.UUID  `json:"parking_lot_id"`
	ParkingSlotId uuid.UUID  `json:"parking_slot_id"`
	TimeFrameId   uuid.UUID  `json:"time_frame_id"`
}

func (ltt *LongTermTicket) TableName() string {
	return "long_term_ticket"
}
