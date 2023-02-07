package route

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
	"parkar-server/conf"
	"parkar-server/pkg/handlers"
	"parkar-server/pkg/repo"
	service2 "parkar-server/pkg/service"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting *extraSetting
}

func NewService() *Service {
	s := &Service{
		service.NewApp("Parkar", "v1.0"),
		&extraSetting{},
	}
	// repo
	_ = env.Parse(s.setting)
	s.Config.DB.DSN = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s connect_timeout=5",
		conf.GetConfig().DBHost,
		conf.GetConfig().DBPort,
		conf.GetConfig().DBUser,
		conf.GetConfig().DBName,
		conf.GetConfig().DBPass,
	)
	db := s.GetDB()
	if s.setting.DbDebugEnable {
		db = db.Debug()
	}
	repoPG := repo.NewPGRepo(db)

	//service
	authService := service2.NewAuthService(repoPG)
	favoriteService := service2.NewFavoriteService(repoPG)
	lotService := service2.NewParkingLotService(repoPG)
	blockService := service2.NewBlockService(repoPG)
	slotService := service2.NewParkingSlotService(repoPG)
	vehicleService := service2.NewVehicleService(repoPG)
	userService := service2.NewUserService(repoPG)
	timeFrameService := service2.NewTimeFrameService(repoPG)

	//handler
	authHandler := handlers.NewAuthHandler(authService)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)
	lotHandler := handlers.NewParkingLotHandler(lotService)
	blockHandler := handlers.NewBlockHandler(blockService)
	slotHandler := handlers.NewParkingSlotHandler(slotService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)
	userHandler := handlers.NewUserHandler(userService)
	timeFrameHandler := handlers.NewTimeFrameHandler(timeFrameService)

	v1Api := s.Router.Group("/api/v1")
	swaggerApi := s.Router.Group("/")

	// swagger
	swaggerApi.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	//auth
	v1Api.POST("/user/login", ginext.WrapHandler(authHandler.Login))
	v1Api.POST("/user/reset-password", ginext.WrapHandler(authHandler.ResetPassword))
	//v1Api.POST("/user/create", ginext.WrapHandler(userHandler.))

	//user
	v1Api.GET("/user/:id", ginext.WrapHandler(userHandler.GetOneUserById))
	v1Api.POST("/user/create", ginext.WrapHandler(userHandler.CreateUser))
	v1Api.POST("/user/check-phone", ginext.WrapHandler(userHandler.CheckDuplicatePhone))
	v1Api.PUT("/user/:id", ginext.WrapHandler(userHandler.UpdateUser))
	v1Api.DELETE("/user/:id", ginext.WrapHandler(userHandler.DeleteUser))

	//favorite
	v1Api.POST("/favorite/create", ginext.WrapHandler(favoriteHandler.Create))
	v1Api.GET("/favorite/user/:idUser", ginext.WrapHandler(favoriteHandler.GetAllFavoriteParkingByUser))
	v1Api.DELETE("/favorite/delete/:id", ginext.WrapHandler(favoriteHandler.DeleteOne))

	//time frame
	v1Api.GET("/time-frame/get-all", ginext.WrapHandler(timeFrameHandler.GetAllTimeFrame))
	v1Api.POST("/time-frame/create-multi", ginext.WrapHandler(timeFrameHandler.Create))
	v1Api.PUT("/time-frame/update", ginext.WrapHandler(timeFrameHandler.Update))

	// parking lot
	v1Api.POST("/parking-lot/create", ginext.WrapHandler(lotHandler.CreateParkingLot))
	v1Api.GET("/parking-lot/get-one/:id", ginext.WrapHandler(lotHandler.GetOneParkingLot))
	v1Api.GET("/parking-lot/get-list", ginext.WrapHandler(lotHandler.GetListParkingLot))
	v1Api.PUT("/parking-lot/update/:id", ginext.WrapHandler(lotHandler.UpdateParkingLot))
	v1Api.DELETE("/parking-lot/delete/:id", ginext.WrapHandler(lotHandler.DeleteParkingLot))

	// block
	v1Api.POST("/block/create", ginext.WrapHandler(blockHandler.CreateBlock))
	v1Api.GET("/block/get-one/:id", ginext.WrapHandler(blockHandler.GetOneBlock))
	v1Api.GET("/block/get-list", ginext.WrapHandler(blockHandler.GetListBlock))
	v1Api.PUT("/block/update/:id", ginext.WrapHandler(blockHandler.UpdateBlock))
	v1Api.DELETE("/block/delete/:id", ginext.WrapHandler(blockHandler.DeleteBlock))

	// parking slot
	v1Api.POST("/parking-slot/create", ginext.WrapHandler(slotHandler.CreateParkingSlot))
	v1Api.GET("/parking-slot/get-one/:id", ginext.WrapHandler(slotHandler.GetOneParkingSlot))
	v1Api.GET("/parking-slot/get-list", ginext.WrapHandler(slotHandler.GetListParkingSlot))
	v1Api.PUT("/parking-slot/update/:id", ginext.WrapHandler(slotHandler.UpdateParkingSlot))
	v1Api.DELETE("/parking-slot/delete/:id", ginext.WrapHandler(slotHandler.DeleteParkingSlot))
	//v1Api.GET("/parking-slot/availability", ginext.WrapHandler(slotHandler.DeleteParkingSlot))

	// parking lot
	v1Api.POST("/vehicle/create", ginext.WrapHandler(vehicleHandler.CreateVehicle))
	v1Api.GET("/vehicle/get-one/:id", ginext.WrapHandler(vehicleHandler.GetOneVehicle))
	v1Api.GET("/vehicle/get-list", ginext.WrapHandler(vehicleHandler.GetListVehicle))
	v1Api.PUT("/vehicle/update/:id", ginext.WrapHandler(vehicleHandler.UpdateVehicle))
	v1Api.DELETE("/vehicle/delete/:id", ginext.WrapHandler(vehicleHandler.DeleteVehicle))

	// Migrate
	migrateHandler := handlers.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrateHandler.Migrate)
	return s
}
