package service

import (
	"context"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
)

type VehicleService struct {
	repo repo.PGInterface
}

func NewVehicleService(repo repo.PGInterface) VehicleServiceInterface {
	return &VehicleService{repo: repo}
}

type VehicleServiceInterface interface {
	CreateTicket(ctx context.Context, ticket *model.Ticket) (*model, error)
	GetAllTicket(ctx context.Context, req model.GetListTicketParam) ([]model.Ticket, error)
}

func (s *VehicleService) CreateTicket(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {
	err := s.repo.CreateTicket(ctx, ticket, nil)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
func (s *VehicleService) GetAllTicket(ctx context.Context, req model.GetListTicketParam) ([]model.Ticket, error) {
	res, err := s.repo.GetAllTicket(ctx, req, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
