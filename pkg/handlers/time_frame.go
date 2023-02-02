package handlers

import (
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/service"
	"parkar-server/pkg/utils"
)

type TimeFrameHandler struct {
	service service.TimeFrameServiceInterface
}

func NewTimeFrameHandler(service service.TimeFrameServiceInterface) TimeFrameHandlerInterface {
	return &TimeFrameHandler{service: service}
}

type TimeFrameHandlerInterface interface {
	GetAllTimeFrame(r *ginext.Request) (*ginext.Response, error)
	Create(r *ginext.Request) (*ginext.Response, error)
	Update(r *ginext.Request) (*ginext.Response, error)
}

func (h *TimeFrameHandler) GetAllTimeFrame(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.GetListTimeFrameParam{}
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	//check valid
	if err := utils.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("Invalid data!")
		return nil, ginext.NewError(http.StatusBadRequest, "Invalid data: "+err.Error())
	}
	res, err := h.service.GetAllTimeFrame(r.GinCtx, req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, res), nil
}

func (h *TimeFrameHandler) Create(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.ListTimeFrameReq{}
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	if len(req.Data) <= 0 {
		log.Error("Số lượng khung giờ phải lớn hơn 0")
		return nil, ginext.NewError(http.StatusBadRequest, "Số lượng khung giờ phải lớn hơn 0")
	}
	err := h.service.CreateTimeFrame(r.GinCtx, req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, "Successfully!"), nil
}
func (h *TimeFrameHandler) Update(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	req := model.ListTimeFrameReq{}
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	if len(req.Data) <= 0 {
		log.Error("Số lượng khung giờ phải lớn hơn 0")
		return nil, ginext.NewError(http.StatusBadRequest, "Số lượng khung giờ phải lớn hơn 0")
	}
	err := h.service.UpdateTimeFrame(r.GinCtx, req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, "Successfully!"), nil
}
