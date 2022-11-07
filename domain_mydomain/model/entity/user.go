package entity

import "demo-cronjobmonitor/domain_mydomain/model/errorenum"

type User struct {
	ID string `` //
	Shipments string
	Shipments_log string
}

type UserRequest struct {
	Shipments string
	Shipments_log string
}

func NewUser(req UserRequest) (*User, error) {
	
	if req.Shipments == "" {
		return nil, errorenum.ShipmentsMustNotEmpty
	}
	if req.Shipments_log == "" {
		return nil, errorenum.Shipments_logMustNotEmpty
	}

	var obj User
	obj.Shipments = req.Shipments
	obj.Shipments_log = req.Shipments_log

	// assign value here

	err := obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *User) Validate() error {
	return nil
}
