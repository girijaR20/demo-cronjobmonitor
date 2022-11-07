package inmemory

import (
	"context"
	"demo-cronjobmonitor/domain_mydomain/model/entity"
	"demo-cronjobmonitor/shared/driver"
	"demo-cronjobmonitor/shared/infrastructure/config"
	"demo-cronjobmonitor/shared/infrastructure/logger"
)

type gateway struct {
	Users   []*entity.User
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) SaveUser(ctx context.Context, obj *entity.User) error {
	r.log.Info(ctx, "called")

	r.Users = append(r.Users)

	return nil
}
