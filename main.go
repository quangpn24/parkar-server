package main

import (
	"context"
	"gitlab.com/goxp/cloud0/logger"
	"os"
	"parkar-server/conf"
	"parkar-server/pkg/route"
	"parkar-server/pkg/utils"
)

const (
	APPNAME = "Parkar"
)

// @title Parkar API
// @version 1.0
// @description This is MVT Adapter api docs.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8088
// @BasePath  /

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()

	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
