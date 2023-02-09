package model

import (
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
	"time"
)

type ParkingLot struct {
	BaseModel
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Lat         string    `json:"lat"`
	Long        string    `json:"long"`
	CompanyID   uuid.UUID `json:"companyID" gorm:"type:uuid"`
}

func (ParkingLot) TableName() string {
	return "parking_lot"
}

type ParkingLotReq struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"name" valid:"Required"`
	Description *string    `json:"description"`
	Address     *string    `json:"address"`
	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	Lat         *string    `json:"lat"`
	Long        *string    `json:"long"`
	CompanyID   *uuid.UUID `json:"companyID"`
}

type ListParkingLotReq struct {
	Name     *string `json:"name" form:"name"`
	Lat      *string `json:"lat" form:"lat"`
	Long     *string `json:"long" form:"long"`
	Sort     string  `json:"sort" form:"sort"`
	Page     int     `json:"page" form:"page"`
	PageSize int     `json:"pageSize" form:"pageSize"`
}

type ListParkingLotRes struct {
	Data []ParkingLot    `json:"data,omitempty"`
	Meta ginext.BodyMeta `json:"meta" swaggertype:"object"`
}
