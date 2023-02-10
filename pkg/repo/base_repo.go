package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math"
	"net/http"
	"parkar-server/pkg/model"
	"parkar-server/pkg/utils"
	"runtime/debug"
	"time"

	"gitlab.com/goxp/cloud0/logger"

	"gitlab.com/goxp/cloud0/ginext"
	"gorm.io/gorm"
)

const (
	StateNew byte = iota + 1 // starts from 1
	StateDoing
	StateDone

	generalQueryTimeout         = 60 * time.Second
	generalQueryTimeout2Minutes = 120 * time.Second
	defaultPageSize             = 30
	maxPageSize                 = 1000
)

func NewPGRepo(db *gorm.DB) PGInterface {
	return &RepoPG{db: db}
}

type PGInterface interface {
	// DB
	DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc)
	DB() (db *gorm.DB)
	Transaction(ctx context.Context, f func(rp PGInterface) error) error

	//user
	GetOneUserByPhone(ctx context.Context, phoneNumber string, tx *gorm.DB) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User, tx *gorm.DB) error
	GetOneUserById(ctx context.Context, id uuid.UUID, tx *gorm.DB) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, tx *gorm.DB) error
	DeleteUser(ctx context.Context, id string, tx *gorm.DB) error

	//favorite
	GetAllFavoriteParkingByUser(ctx context.Context, userId string, tx *gorm.DB) (res []model.Favorite, err error)
	CreateFavorite(ctx context.Context, favorite *model.Favorite, tx *gorm.DB) error
	DeleteOneFavorite(ctx context.Context, id uuid.UUID, tx *gorm.DB) error
	GetOne(ctx context.Context, req model.FavoriteRequestV2, tx *gorm.DB) (model.Favorite, error)

	//time frame
	GetAllTimeFrame(ctx context.Context, req model.GetListTimeFrameParam, tx *gorm.DB) (res *model.ListTimeFrame, err error)
	CreateMultiTimeFrame(ctx context.Context, timeFrame []model.TimeFrame, tx *gorm.DB) (err error)
	DeleteTimeFrameByParkingLotID(ctx context.Context, parkingLotID string, tx *gorm.DB) (err error)
	//ticket
	CreateTicket(ctx context.Context, req *model.Ticket, tx *gorm.DB) error
	GetAllTicket(ctx context.Context, req model.GetListTicketParam, tx *gorm.DB) ([]model.Ticket, error)
	GetOneTicket(ctx context.Context, id string, tx *gorm.DB) (model.Ticket, error)
	GetOneTicketWithExtend(ctx context.Context, id string, tx *gorm.DB) (model.Ticket, error)
	GetListExtendTicketByOrigin(ctx context.Context, idParent string, tx *gorm.DB) ([]model.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *model.Ticket, tx *gorm.DB) error
	GetAllTicketCompany(ctx context.Context, req model.GetListTicketReq) (res []model.GetListTicketRes, err error)

	//ticket extend
	CreateTicketExtend(ctx context.Context, req *model.TicketExtend, tx *gorm.DB) error
	//long term ticket
	CreateLongTermTicket(ctx context.Context, ltTicket *model.LongTermTicket, tx *gorm.DB) error

	//token
	CreateRefreshToken(ctx context.Context, refreshToken *model.RefreshToken, tx *gorm.DB) error

	// Parking lot
	CreateParkingLot(ctx context.Context, req *model.ParkingLot) error
	GetOneParkingLot(ctx context.Context, id uuid.UUID) (model.ParkingLot, error)
	GetListParkingLot(ctx context.Context, req model.ListParkingLotReq) (model.ListParkingLotRes, error)
	GetListParkingLotCompany(ctx context.Context, req model.GetListParkingLotReq) (model.ListParkingLotRes, error)
	UpdateParkingLot(ctx context.Context, req *model.ParkingLot) error
	DeleteParkingLot(ctx context.Context, id uuid.UUID) error

	// Block
	CreateBlock(ctx context.Context, req *model.Block) error
	GetOneBlock(ctx context.Context, id uuid.UUID) (model.Block, error)
	GetListBlock(ctx context.Context, req model.ListBlockReq) (model.ListBlockRes, error)
	UpdateBlock(ctx context.Context, req *model.Block) error
	DeleteBlock(ctx context.Context, id uuid.UUID) error

	// ParkingSlot
	CreateParkingSlot(ctx context.Context, req *model.ParkingSlot) error
	GetOneParkingSlot(ctx context.Context, id uuid.UUID) (model.ParkingSlot, error)
	GetListParkingSlot(ctx context.Context, req model.ListParkingSlotReq) (model.ListParkingSlotRes, error)
	GetAvailableParkingSlot(ctx context.Context, req model.AvailableParkingSlotReq) (model.ListParkingSlotRes, error)
	UpdateParkingSlot(ctx context.Context, req *model.ParkingSlot) error
	DeleteParkingSlot(ctx context.Context, id uuid.UUID) error

	// Vehicle
	CreateVehicle(ctx context.Context, req *model.Vehicle) error
	GetOneVehicle(ctx context.Context, id uuid.UUID) (model.Vehicle, error)
	GetListVehicle(ctx context.Context, req model.ListVehicleReq) (model.ListVehicleRes, error)
	UpdateVehicle(ctx context.Context, req *model.Vehicle) error
	DeleteVehicle(ctx context.Context, id uuid.UUID) error

	// company
	CreateCompany(ctx context.Context, req *model.Company) error
	GetCompanyByEmail(ctx context.Context, email string) (model.Company, error)
	GetOneCompany(ctx context.Context, id uuid.UUID) (model.Company, error)
	UpdateCompany(ctx context.Context, req *model.Company) error
}

type RepoPG struct {
	db    *gorm.DB
	debug bool
}

func (r *RepoPG) Transaction(ctx context.Context, f func(rp PGInterface) error) (err error) {
	log := logger.WithCtx(ctx, "RepoPG.Transaction")
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	// create new instance to run the transaction
	repo := *r
	tx = tx.Begin()
	repo.db = tx
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.New(fmt.Sprint(r))
			log.WithError(err).Error("error_500: Panic when run Transaction")
			debug.PrintStack()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	err = f(&repo)
	if err != nil {
		log.WithError(err).Error("error_500: Error when run Transaction")
		return err
	}
	return nil
}

func (r *RepoPG) DB() *gorm.DB {
	return r.db
}

func (r *RepoPG) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func (r *RepoPG) DBWithTimeout2Minutes(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout2Minutes)
	return r.db.WithContext(ctx), cancel
}

func (r *RepoPG) GetPage(page int) int {
	if page == 0 {
		return 1
	}
	return page
}

func (r *RepoPG) GetOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

func (r *RepoPG) GetPageSize(pageSize int) int {
	if pageSize == 0 {
		return defaultPageSize
	}
	if pageSize > maxPageSize {
		return maxPageSize
	}
	return pageSize
}

func (r *RepoPG) GetTotalPages(totalRows, pageSize int) int {
	return int(math.Ceil(float64(totalRows) / float64(pageSize)))
}

func (r *RepoPG) GetOrder(sort string) string {
	if sort == "" {
		sort = "created_at desc"
	}
	return sort
}

func (r *RepoPG) GetOrderBy(sort string) string {
	if sort == "" {
		sort = "created_at desc"
	}
	return sort
}

func (r *RepoPG) GetPaginationInfo(query string, tx *gorm.DB, totalRow, page, pageSize int) (rs ginext.BodyMeta, err error) {
	tm := struct {
		Count int `json:"count"`
	}{}
	if query != "" {
		if err = tx.Raw(query).Scan(&tm).Error; err != nil {
			return nil, err
		}
		totalRow = tm.Count
	}

	return ginext.BodyMeta{
		"page":        page,
		"page_size":   pageSize,
		"total_pages": r.GetTotalPages(totalRow, pageSize),
		"total_rows":  totalRow,
	}, nil
}

func (r *RepoPG) ReturnErrorInGetFuncV2(ctx context.Context, logStr string, err error, key string, value interface{}) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0)).WithField(key, value)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.WithError(err).Error(fmt.Sprintf("error_404: %s - RepoPG", logStr))
		return ginext.NewError(http.StatusNotFound, utils.MessageError()[http.StatusNotFound])
	}
	log.WithError(err).Error(fmt.Sprintf("error_500: %s - RepoPG", logStr))
	return ginext.NewError(http.StatusInternalServerError, err.Error())
}
