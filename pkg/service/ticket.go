package service

import (
	"context"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
)

type TicketService struct {
	repo repo.PGInterface
}

func NewTicketService(repo repo.PGInterface) TicketServiceInterface {
	return &TicketService{repo: repo}
}

type TicketServiceInterface interface {
	CreateTicket(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error)
	GetAllTicket(ctx context.Context, req model.GetListTicketParam) ([]model.Ticket, error)
	CancelTicket(ctx context.Context, req model.CancelTicketRequest) error
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {
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
