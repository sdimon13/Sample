package entity

import (
	"git.sample.ru/sample/pkg/golibs/nulltype"
	"time"
)

type Service struct {
	Id        int64             `json:"id"`
	MasterId  int64             `json:"master_id"`
	Name      string            `json:"name"`
	Price     int32             `json:"price"`
	Time      time.Time         `json:"time"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt nulltype.NullTime `json:"deleted_at"`
}
