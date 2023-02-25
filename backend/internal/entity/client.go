package entity

import (
	"git.sample.ru/sample/pkg/golibs/nulltype"
	"time"
)

type Client struct {
	Id        int64             `json:"id"`
	Phone     string            `json:"phone"`
	Name      string            `json:"name"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt nulltype.NullTime `json:"deleted_at"`
}

func (e Client) GetTable() string {
	return "samples.clients"
}
