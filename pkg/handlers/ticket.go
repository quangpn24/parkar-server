package handlers

import (
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/service"
	"parkar-server/pkg/utils"
	"parkar-server/pkg/valid"
)

type TicketHandler struct {
	service service.TicketServiceInterface
}

func NewTicketHandler(service service.TicketServiceInterface) TicketHandlerInterface {
	return &TicketHandler{service: service}
}

type TicketHandlerInterface interface {
	CreateTicket(r *ginext.Request) (*ginext.Response, error)
	GetAllTicket(r *ginext.Request) (*ginext.Response, error)
	GetOneTicketWithExtend(r *ginext.Request) (*ginext.Response, error)
	CancelTicket(r *ginext.Request) (*ginext.Response, error)
	ExtendTicket(r *ginext.Request) (*ginext.Response, error)
}

func (h *TicketHandler) CreateTicket(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))
	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}
	req := model.TicketReq{}
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	res, err := h.service.CreateTicket(r.Context(), &req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusCreated, res), nil
}
func (h *TicketHandler) ExtendTicket(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))
	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}
	req := model.ExtendTicketReq{}
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	res, err := h.service.ExtendTicket(r.Context(), &req)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusCreated, res), nil
}
func (h *TicketHandler) GetAllTicket(r *ginext.Request) (*ginext.Response, error) {
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
func (h *TicketHandler) GetOneTicketWithExtend(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))

	ticketId := utils.ParseIDFromUri(r.GinCtx)

	if ticketId == nil {
		log.Error("Ticket id is required")
		return nil, ginext.NewError(http.StatusBadRequest, "Ticket id is required")
	}
	res, err := h.service.GetOneTicketWithExtend(r.Context(), valid.UUID(ticketId).String())
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, res), nil
}
func (h *TicketHandler) CancelTicket(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, utils.GetCurrentCaller(h, 0))
	req := model.CancelTicketRequest{}
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("Error when parse req!")
		return nil, ginext.NewError(http.StatusBadRequest, "Error when parse req: "+err.Error())
	}
	err := h.service.CancelTicket(r.Context(), req.TicketId)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponse(http.StatusOK), nil
}
