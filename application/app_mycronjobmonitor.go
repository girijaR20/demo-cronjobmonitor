package application

import (
	"demo-cronjobmonitor/domain_mydomain/controller/userapi"
	"demo-cronjobmonitor/domain_mydomain/gateway/inmemory"
	"demo-cronjobmonitor/domain_mydomain/usecase/cronjobmonitor"
	"demo-cronjobmonitor/shared/driver"
	"demo-cronjobmonitor/shared/infrastructure/config"
	"demo-cronjobmonitor/shared/infrastructure/logger"
	"demo-cronjobmonitor/shared/infrastructure/server"
	"demo-cronjobmonitor/shared/util"
)

type myCronJobMonitor struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c myCronJobMonitor) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewMyCronJobMonitor() func() driver.RegistryContract {

	const appName = "myCronJobMonitor"

	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData(appName, appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

		datasource := inmemory.NewGateway(log, appData, cfg)

		return &myCronJobMonitor{
			httpHandler: &httpHandler,
			controller: &userapi.Controller{
				Log:                  log,
				Config:               cfg,
				Router:               httpHandler.Router,
				CronJobMonitorInport: cronjobmonitor.NewUsecase(datasource),
			},
		}

	}
}
