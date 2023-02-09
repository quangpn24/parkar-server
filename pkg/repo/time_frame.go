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

func (r *RepoPG) GetAllTimeFrame(ctx context.Context, req model.GetListTimeFrameParam, tx *gorm.DB) (*model.ListTimeFrame, error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	res := &model.ListTimeFrame{}
	if err := tx.Model(&model.TimeFrame{}).Where("parking_lot_id = ? ", req.ParkingLotId).Find(&res.Data).Error; err != nil {
		log.WithError(err).Error("Error when get all time frame - GetAllTimeFrame - RepoPG")
		return nil, ginext.NewError(http.StatusInternalServerError, "Error when get all time frame: "+err.Error())
	}
	return res, nil
}
func (r *RepoPG) CreateMultiTimeFrame(ctx context.Context, timeFrame []model.TimeFrame, tx *gorm.DB) (err error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	if err := tx.Model(&model.TimeFrame{}).Create(&timeFrame).Error; err != nil {
		log.WithError(err).Error("Error when create time frame - CreateMultiTimeFrame - RepoPG")
		return ginext.NewError(http.StatusInternalServerError, "Error when create time frame: "+err.Error())
	}
	return nil
}
func (r *RepoPG) DeleteTimeFrameByParkingLotID(ctx context.Context, parkingLotID string, tx *gorm.DB) (err error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err := tx.Where("parking_lot_id = ?", parkingLotID).Delete(&model.TimeFrame{}).Error; err != nil {
		log.WithError(err).Error("error_500: error when delete time frame - DeleteTimeFrameByParkingLotID -RepoPG")
		return ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
