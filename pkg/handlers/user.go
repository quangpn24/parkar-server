package handlers

import (
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/service"
	"parkar-server/pkg/utils"
)

type UserHandler struct {
	service service.TicketServiceInterface
}

func NewUserHandler(service service.TicketServiceInterface) UserHandlerInterface {
	return &UserHandler{service: service}
}

type UserHandlerInterface interface {
	GetAllUser(r *ginext.Request) (*ginext.Response, error)
}

func (h *UserHandler) GetAllUser(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))
	req := model.GetListTicketParam{}
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	//check valid
	if err := utils.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("Invalid data!")
		return nil, ginext.NewError(http.StatusBadRequest, "Invalid data: "+err.Error())
	}
	res, err := h.service.GetAllTicket(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, res), nil
}
