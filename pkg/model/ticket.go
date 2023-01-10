package model

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	BaseModel
	VehicleId        uuid.UUID `json:"vehicle_id"`
	ParkingLotId     uuid.UUID `json:"parking_lot_id"`
	ParkingSlotId    uuid.UUID `json:"parking_slot_id"`
	TimeFrameId      uuid.UUID `json:"time_frame_id"`
	TimeFrame        *TimeFrame
	StartTime        *time.Time `json:"start_time"`
	EndTime          *time.Time `json:"end_time"`
	EntryTime        *time.Time `json:"entry_time"`
	ExitTime         *time.Time `json:"exit_time"`
	Total            float64    `json:"total"`
	State            string     `json:"state"`
	IsExtend         bool       `json:"is_extend"`
	LongTermTicketId uuid.UUID  `json:"long_term_ticket_id"`
}

func (t *Ticket) TableName() string {
	return "ticket"
}
