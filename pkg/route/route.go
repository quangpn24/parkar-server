package route

import (
	"github.com/caarlos0/env/v6"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
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
	db := s.GetDB()
	if s.setting.DbDebugEnable {
		db = db.Debug()
	}
	repoPG := repo.NewPGRepo(db)

	//service
	authService := service2.NewAuthService(repoPG)
	favoriteService := service2.NewFavoriteService(repoPG)

	//handler
	authHandler := handlers.NewAuthHandler(authService)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)

	v1Api := s.Router.Group("/api/v1")
	swaggerApi := s.Router.Group("/")

	// swagger
	swaggerApi.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	//auth
	v1Api.POST("/user/login", ginext.WrapHandler(authHandler.Login))

	//favorite
	v1Api.POST("/favorite", ginext.WrapHandler(favoriteHandler.Create))
	v1Api.GET("/favorite/user/:idUser", ginext.WrapHandler(favoriteHandler.GetAllFavoriteParkingByUser))
	v1Api.DELETE("/favorite", ginext.WrapHandler(favoriteHandler.DeleteOne))

	// Migrate
	migrateHandler := handlers.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrateHandler.Migrate)
	return s
}
