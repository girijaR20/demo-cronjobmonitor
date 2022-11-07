package cronjobmonitor

import "demo-cronjobmonitor/domain_mydomain/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveUserRepo
}
