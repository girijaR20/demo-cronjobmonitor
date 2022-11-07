package errorenum

import "demo-cronjobmonitor/shared/model/apperror"

const (
	SomethingError            apperror.ErrorType = "ER0000 something error"
	//CronJobMonitor            apperror.ErrorType = "ER0001 cron job monitor"
	ShipmentsMustNotEmpty     apperror.ErrorType = "ER0002 shipments must not empty"
	Shipments_logMustNotEmpty apperror.ErrorType = "ER0003 shipments_log must not empty"
)
