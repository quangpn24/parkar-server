package repo

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"gorm.io/gorm"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/utils"
)

func (r *RepoPG) GetOneUserById(ctx context.Context, id uuid.UUID, tx *gorm.DB) (*model.User, error) {
	var cancel context.CancelFunc
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	rs := &model.User{}

	if err := tx.Model(&model.User{}).Where("id = ?", id.String()).Take(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.WithError(err).Error("Not found user by id: " + id.String())
			return nil, ginext.NewError(http.StatusInternalServerError, "Not found user by id: "+id.String())
		} else {
			log.WithError(err).Error("Error when get one user by id")
			return nil, ginext.NewError(http.StatusInternalServerError, "Error when get one user by id")
		}
	}
	return rs, nil
}
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
func (r *RepoPG) UpdateUser(ctx context.Context, user *model.User, tx *gorm.DB) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err := tx.Model(&model.User{}).Where("id = ?", user.ID).Save(&user).Error; err != nil {
		log.WithError(err).Error("error_500: error when UpdateUser")
		return ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
func (r *RepoPG) DeleteUser(ctx context.Context, id string, tx *gorm.DB) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	if err := tx.Where("id = ? ", id).Delete(&model.User{}).Error; err != nil {
		log.WithError(err).Error("Error when delete user - DeleteUser - RepoPG")
		return ginext.NewError(http.StatusInternalServerError, "Error when delete user: "+err.Error())
	}
	return nil
}
