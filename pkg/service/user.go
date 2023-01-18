package service

import (
	"context"
	"github.com/google/uuid"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
	"parkar-server/pkg/utils"
	"parkar-server/pkg/valid"
)

type UserService struct {
	repo repo.PGInterface
}

func NewUserService(repo repo.PGInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

type UserServiceInterface interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
	CheckDuplicatePhone(ctx context.Context, phoneNumber string) (bool, error)
	UpdateUser(ctx context.Context, userReq model.UserReq) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func (s *UserService) GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	rs, err := s.repo.GetOneUserById(ctx, id, nil)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
func (s *UserService) CheckDuplicatePhone(ctx context.Context, phone string) (bool, error) {
	rs, err := s.repo.GetOneUserByPhone(ctx, phone, nil)
	if err != nil {
		return false, err
	}
	if rs != nil {
		return true, nil
	}
	return false, nil
}
func (s *UserService) UpdateUser(ctx context.Context, userReq model.UserReq) (*model.User, error) {
	user, err := s.repo.GetOneUserById(ctx, valid.UUID(userReq.ID), nil)
	if err != nil {
		return nil, err
	}

	utils.Sync(userReq, &user)
	if err := s.repo.UpdateUser(ctx, user, nil); err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id, nil)
}
