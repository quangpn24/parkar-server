package repo

import (
	"context"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"gorm.io/gorm"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/utils"
)

func (r *RepoPG) GetOneUserByPhone(ctx context.Context, phoneNumber string, tx *gorm.DB) (*model.User, error) {
	var cancel context.CancelFunc
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	rs := &model.User{}

	if err := tx.Model(&model.User{}).Where("phone_number = ?", phoneNumber).Take(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.WithError(err).Error("Not found user by phone number: " + phoneNumber)
			return nil, ginext.NewError(http.StatusInternalServerError, "Not found user by phone number: "+phoneNumber)
		} else {
			log.WithError(err).Error("Error when get one user by phone number")
			return nil, ginext.NewError(http.StatusInternalServerError, "Error when get one user by phone number")
		}
	}
	return rs, nil
}
func (r *RepoPG) CreateUser(ctx context.Context, user *model.User, tx *gorm.DB) error {
	var cancel context.CancelFunc
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err := tx.Model(&model.User{}).Create(&user).Error; err != nil {
		log.WithError(err).Error("Error when create user")
		return ginext.NewError(http.StatusInternalServerError, "Error when create user: "+err.Error())
	}
	return nil
}

func (r *RepoPG) CreateRefreshToken(ctx context.Context, refreshToken *model.RefreshToken, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	if err := tx.Create(&refreshToken).Error; err != nil {
		return ginext.NewError(http.StatusInternalServerError, "Error when create refresh token:"+err.Error())
	}
	return nil
}
