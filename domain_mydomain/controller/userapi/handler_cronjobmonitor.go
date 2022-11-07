package userapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"demo-cronjobmonitor/domain_mydomain/usecase/cronjobmonitor"
	"demo-cronjobmonitor/shared/infrastructure/logger"
	"demo-cronjobmonitor/shared/model/payload"
	"demo-cronjobmonitor/shared/util"
)

// cronJobMonitorHandler ...
func (r *Controller) cronJobMonitorHandler() gin.HandlerFunc {

	type request struct {
		Shipments     string `json:"shipments"`
		Shipments_log string `json:"shipments_log"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req cronjobmonitor.InportRequest
		req.Shipments = jsonReq.Shipments
		req.Shipments_log = jsonReq.Shipments_log

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.CronJobMonitorInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
