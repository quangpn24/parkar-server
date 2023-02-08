package service

import (
	"context"
	"github.com/google/uuid"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
	"parkar-server/pkg/utils"
	"parkar-server/pkg/valid"
)

type ParkingSlotService struct {
	repo repo.PGInterface
}

func NewParkingSlotService(repo repo.PGInterface) ParkingSlotInterface {
	return &ParkingSlotService{repo: repo}
}

type ParkingSlotInterface interface {
	CreateParkingSlot(ctx context.Context, req model.ParkingSlotReq) (*model.ParkingSlot, error)
	GetListParkingSlot(ctx context.Context, req model.ListParkingSlotReq) (model.ListParkingSlotRes, error)
	GetAvailableParkingSlot(ctx context.Context, req model.AvailableParkingSlotReq) (model.ListBlockRes, error)
	GetOneParkingSlot(ctx context.Context, id uuid.UUID) (model.ParkingSlot, error)
	UpdateParkingSlot(ctx context.Context, req model.ParkingSlotReq) (model.ParkingSlot, error)
	DeleteParkingSlot(ctx context.Context, id uuid.UUID) error
}

func (s *ParkingSlotService) CreateParkingSlot(ctx context.Context, req model.ParkingSlotReq) (*model.ParkingSlot, error) {
	ParkingSlot := &model.ParkingSlot{
		Name:        valid.String(req.Name),
		Description: valid.String(req.Description),
		BlockID:     valid.UUID(req.BlockID),
	}

	if err := s.repo.CreateParkingSlot(ctx, ParkingSlot); err != nil {
		return nil, err
	}
	return ParkingSlot, nil
}

func (s *ParkingSlotService) GetListParkingSlot(ctx context.Context, req model.ListParkingSlotReq) (model.ListParkingSlotRes, error) {
	return s.repo.GetListParkingSlot(ctx, req)
}
func (s *ParkingSlotService) GetAvailableParkingSlot(ctx context.Context, req model.AvailableParkingSlotReq) (model.ListBlockRes, error) {
	res, err := s.repo.GetAvailableParkingSlot(ctx, req)
	if err != nil {
		return model.ListBlockRes{}, err
	}
	blockMap := map[string]model.Block{}
	result := model.ListBlockRes{}
	for _, slot := range res.Data {
		if val, ok := blockMap[slot.BlockID.String()]; ok {
			slot.Block = nil
			val.ParkingSLots = append(val.ParkingSLots, slot)
			blockMap[slot.BlockID.String()] = val
		} else {
			newBlock := slot.Block
			slot.Block = nil
			newBlock.ParkingSLots = append(newBlock.ParkingSLots, slot)
			blockMap[slot.BlockID.String()] = *newBlock
		}
	}
	for _, val := range blockMap {
		result.Data = append(result.Data, val)
	}
	return result, nil
}

func (s *ParkingSlotService) GetOneParkingSlot(ctx context.Context, id uuid.UUID) (model.ParkingSlot, error) {
	return s.repo.GetOneParkingSlot(ctx, id)
}

func (s *ParkingSlotService) UpdateParkingSlot(ctx context.Context, req model.ParkingSlotReq) (model.ParkingSlot, error) {
	ParkingSlot, err := s.repo.GetOneParkingSlot(ctx, valid.UUID(req.ID))
	if err != nil {
		return ParkingSlot, err
	}

	utils.Sync(req, &ParkingSlot)
	if err := s.repo.UpdateParkingSlot(ctx, &ParkingSlot); err != nil {
		return ParkingSlot, err
	}

	return ParkingSlot, nil
}

func (s *ParkingSlotService) DeleteParkingSlot(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteParkingSlot(ctx, id)
}
