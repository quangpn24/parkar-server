package service

import (
	"context"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
	"parkar-server/pkg/valid"
)

type TicketService struct {
	repo repo.PGInterface
}

func NewTicketService(repo repo.PGInterface) TicketServiceInterface {
	return &TicketService{repo: repo}
}

type TicketServiceInterface interface {
	CreateTicket(ctx context.Context, req *model.TicketReq) (*model.Ticket, error)
	GetAllTicket(ctx context.Context, req model.GetListTicketParam) ([]model.Ticket, error)
	CancelTicket(ctx context.Context, req model.CancelTicketRequest) error
}

func (s *TicketService) CreateTicket(ctx context.Context, req *model.TicketReq) (*model.Ticket, error) {
	ticket := &model.Ticket{
		BaseModel: model.BaseModel{
			CreatorID: req.UserId,
			UpdaterID: req.UserId,
		},
		UserId:        req.UserId,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		VehicleId:     req.VehicleId,
		ParkingLotId:  req.ParkingLotId,
		ParkingSlotId: req.ParkingSlotId,
		TimeFrameId:   req.TimeFrameId,
		State:         "new",
		Total:         valid.Float64(req.Total),
	}

	if req.IsLongTerm {
		longTermTicket := &model.LongTermTicket{
			BaseModel: model.BaseModel{
				CreatorID: req.UserId,
				UpdaterID: req.UserId,
			},
			StartTime:     req.StartTime,
			EndTime:       req.EndTime,
			VehicleId:     req.VehicleId,
			ParkingLotId:  req.ParkingLotId,
			ParkingSlotId: req.ParkingSlotId,
			TimeFrameId:   req.TimeFrameId,
		}
		s.repo.CreateLongTermTicket(ctx, longTermTicket, nil)

		//create ticket normal
		switch req.Type {
		case "DAILY":
		case "CYCLE":
		case "CUSTOM":
		}
	}
	err := s.repo.CreateTicket(ctx, ticket, nil)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
func (s *TicketService) GetAllTicket(ctx context.Context, req model.GetListTicketParam) ([]model.Ticket, error) {
	res, err := s.repo.GetAllTicket(ctx, req, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *TicketService) CancelTicket(ctx context.Context, req model.CancelTicketRequest) error {
	return s.repo.CancelTicket(ctx, req, nil)
}
