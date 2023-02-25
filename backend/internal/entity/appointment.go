package entity

import (
	"git.sample.ru/sample/pkg/golibs/nulltype"
	"time"
)

type Appointment struct {
	Id              int64               `json:"id"`
	MasterId        int64               `json:"master_id"`
	ClientId        nulltype.NullInt64  `json:"client_id,omitempty"`
	ServiceId       nulltype.NullInt64  `json:"service_id"`
	StatusId        int32               `json:"status_id"`
	AppointmentDate time.Time           `json:"appointment_date"`
	AppointmentTime time.Time           `json:"appointment_time"`
	Note            nulltype.NullString `json:"note,omitempty"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
	DeletedAt       nulltype.NullTime   `json:"deleted_at"`
}

func (e Appointment) GetTable() string {
	return "samples.appointments"
}
