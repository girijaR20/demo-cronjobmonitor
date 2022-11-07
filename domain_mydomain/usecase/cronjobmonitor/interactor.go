package cronjobmonitor

import (
	"context"
	"demo-cronjobmonitor/domain_mydomain/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type cronJobMonitorInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &cronJobMonitorInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *cronJobMonitorInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	userObj, err := entity.NewUser(entity.UserRequest{
		Shipments:     req.Shipments,
		Shipments_log: req.Shipments_log,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}

	//!

	return res, nil
}
