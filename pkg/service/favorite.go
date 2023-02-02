package service

import (
	"context"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
	"parkar-server/pkg/valid"
)

type FavoriteService struct {
	repo repo.PGInterface
}

func NewFavoriteService(repo repo.PGInterface) FavoriteServiceInterface {
	return &FavoriteService{repo: repo}
}

type FavoriteServiceInterface interface {
	GetAllFavoriteParkingByUser(ctx context.Context, userId string) (res []model.Favorite, err error)
	Create(ctx context.Context, req model.FavoriteRequest) (res *model.Favorite, err error)
	DeleteOne(ctx context.Context, req model.FavoriteRequest) error
}

func (s *FavoriteService) Create(ctx context.Context, req model.FavoriteRequest) (res *model.Favorite, err error) {
	favorite := &model.Favorite{
		UserId:       req.UserId,
		ParkingLotId: req.ParkingLotId,
		BaseModel: model.BaseModel{
			CreatorID: valid.UUIDPointer(req.UserId),
			UpdaterID: valid.UUIDPointer(req.UserId),
		},
	}
	if err := s.repo.CreateFavorite(ctx, favorite, nil); err != nil {
		return nil, err
	}
	return favorite, nil
}
func (s *FavoriteService) GetAllFavoriteParkingByUser(ctx context.Context, userId string) (res []model.Favorite, err error) {
	rs, err := s.repo.GetAllFavoriteParkingByUser(ctx, userId, nil)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *FavoriteService) DeleteOne(ctx context.Context, req model.FavoriteRequest) error {
	if err := s.repo.DeleteOneFavorite(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
