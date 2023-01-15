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

func (r *RepoPG) CreateTicket(ctx context.Context, ticket *model.Ticket, tx *gorm.DB) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	if err := tx.Model(&model.Ticket{}).Create(&ticket).Error; err != nil {
		log.WithError(err).Error("Error when create ticket - CreateTicket - RepoPG")
		return ginext.NewError(http.StatusInternalServerError, "Error when create ticket: "+err.Error())
	}
	return nil
}
func (r *RepoPG) CancelTicket(ctx context.Context, req model.CancelTicketRequest, tx *gorm.DB) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	if err := tx.Model(&model.Ticket{}).Where("id in ?", req.ListTicketId).Updates(map[string]interface{}{"state": "cancel"}).Error; err != nil {
		log.WithError(err).Error("Error when cancel ticket - CreateTicket - RepoPG")
		return ginext.NewError(http.StatusInternalServerError, "Error when cancel ticket: "+err.Error())
	}
	return nil
}
func (r *RepoPG) GetAllTicket(ctx context.Context, req model.GetListTicketParam, tx *gorm.DB) (res []model.Ticket, err error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	tx = tx.Model(&model.Ticket{})
	if req.State != nil {
		tx = tx.Where("state = ?", req.State)
	}
	if err := tx.Where("user_id = ?", req.UserId).Find(&res).Error; err != nil {
		log.WithError(err).Error("Error when get all ticket - GetAllTicket - RepoPG")
		return nil, ginext.NewError(http.StatusInternalServerError, "Error when get all ticket: "+err.Error())
	}
	return res, nil
}
