package cronjobmonitor

import (
	"context"

	"demo-cronjobmonitor/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Shipments string
	Shipments_log string
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}
