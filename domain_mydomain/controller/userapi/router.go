package userapi

import (
	"github.com/gin-gonic/gin"

	"demo-cronjobmonitor/domain_mydomain/usecase/cronjobmonitor"
	"demo-cronjobmonitor/shared/infrastructure/config"
	"demo-cronjobmonitor/shared/infrastructure/logger"
)

type Controller struct {
	Router               gin.IRouter
	Config               *config.Config
	Log                  logger.Logger
	CronJobMonitorInport cronjobmonitor.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/cronjobmonitor", r.authorized(), r.cronJobMonitorHandler())
	r.Router.GET("/cronjobmonitor", r.authorized(), r.cronJobMonitorHandler())
}
