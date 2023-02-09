package service

import (
	"context"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/repo"
	"parkar-server/pkg/utils"
	"parkar-server/pkg/valid"
	"time"
)

type AuthService struct {
	repo repo.PGInterface
}

func NewAuthService(repo repo.PGInterface) AuthServiceInterface {
	return &AuthService{repo: repo}
}

type AuthServiceInterface interface {
	Login(ctx context.Context, req model.Credential) (interface{}, error)
	ResetPassword(ctx context.Context, req model.Credential) error
}

func (s *AuthService) Login(ctx context.Context, req model.Credential) (interface{}, error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(s, 0))
	user, err := s.repo.GetOneUserByPhone(ctx, valid.String(req.UserName), nil)
	if err != nil {
		return nil, err
	}

	//check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(valid.String(req.Password))); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, "Mật khẩu không đúng!")
	}

	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		log.WithError(err).Error("Error when generate token - Login - AuthService")
		return nil, err
	}
	//refresh token
	exp := time.Now().Add(utils.EXPIRTE_TIME * time.Second)
	rf_token, _ := utils.GenerateToken(user.ID.String())
	refreshToken := &model.RefreshToken{
		Token:       rf_token,
		ExpiredDate: valid.DayTimePointer(exp),
	}
	if err := s.repo.CreateRefreshToken(ctx, refreshToken, nil); err != nil {
		return nil, err
	}
	res := model.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken.Token,
		PhoneNumber:  user.PhoneNumber,
		DisplayName:  user.DisplayName,
		Id:           user.ID,
		Email:        user.Email,
		ImageUrl:     user.ImageUrl,
	}
	return res, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req model.Credential) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(s, 0))
	user, err := s.repo.GetOneUserByPhone(ctx, valid.String(req.UserName), nil)
	if err != nil {
		return err
	}
	hashPass, err := utils.Hash(valid.String(req.Password))
	if err != nil {
		log.WithError(err).Error("Failed to hash password")
		return ginext.NewError(http.StatusInternalServerError, "Failed to hash password")
	}
	user.Password = hashPass

	if err := s.repo.UpdateUser(ctx, user, nil); err != nil {
		return err
	}
	return nil
}
