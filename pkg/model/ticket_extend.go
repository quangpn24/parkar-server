package model

import "github.com/google/uuid"

type TicketExtend struct {
	BaseModel
	TicketId       uuid.UUID `json:"ticket_id"`
	TicketExtendId uuid.UUID `json:"ticket_extend_id"`
}

func (tx *TicketExtend) TableName() string {
	return "ticket_extend"
}
