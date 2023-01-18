package service

import (
	"context"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
	"parkar-server/pkg/utils"
)

type TimeFrameService struct {
	repo repo.PGInterface
}

func NewTimeFrameService(repo repo.PGInterface) TimeFrameServiceInterface {
	return &TimeFrameService{repo: repo}
}

type TimeFrameServiceInterface interface {
	GetAllTimeFrame(ctx context.Context, req model.GetListTimeFrameParam) (res model.ListTimeFrame, err error)
	CreateTimeFrame(ctx context.Context, req model.ListTimeFrameReq) (err error)
	UpdateTimeFrame(ctx context.Context, req model.ListTimeFrameReq) (err error)
}

func (s *TimeFrameService) GetAllTimeFrame(ctx context.Context, req model.GetListTimeFrameParam) (model.ListTimeFrame, error) {
	res, err := s.repo.GetAllTimeFrame(ctx, req, nil)
	if err != nil {
		return model.ListTimeFrame{}, err
	}
	return *res, nil
}
func (s *TimeFrameService) CreateTimeFrame(ctx context.Context, req model.ListTimeFrameReq) (err error) {
	listUser := []model.TimeFrame{}
	for _, item := range req.Data {
		tmp := model.TimeFrame{}
		utils.Sync(item, &tmp)
		listUser = append(listUser, tmp)
	}
	err = s.repo.CreateMultiTimeFrame(ctx, listUser, nil)
	return err
}
func (s *TimeFrameService) UpdateTimeFrame(ctx context.Context, req model.ListTimeFrameReq) (err error) {
	//detele all time fram by parking lot
	err = s.repo.DeleteTimeFrameByParkingLotID(ctx, req.Data[0].ParkingLotId.String(), nil)
	if err != nil {
		return err
	}
	err = s.CreateTimeFrame(ctx, req)
	return err
}
