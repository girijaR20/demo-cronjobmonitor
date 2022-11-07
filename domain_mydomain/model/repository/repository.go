package repository

import (
	"context"
	"demo-cronjobmonitor/domain_mydomain/model/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}
