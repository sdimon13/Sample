package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullInt64 struct {
	Data sql.NullInt64
}

func NewNullInt64(data *int64) NullInt64 {
	if data == nil {
		return NullInt64{}
	}
	return NullInt64{
		Data: sql.NullInt64{
			Int64: *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *NullInt64) Valid() bool {
	return t.Data.Valid
}
func (t *NullInt64) Get() int64 {
	return t.Data.Int64
}
func (t *NullInt64) GetPtr() *int64 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullInt64) Reset() {
	t.Data = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
}
func (t *NullInt64) Set(data int64) {
	t.Data = sql.NullInt64{
		Int64: data,
		Valid: true,
	}
}

// fmt

func (t *NullInt64) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Int64)
}

// SQL

func (t *NullInt64) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullInt64) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullInt64) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullInt64{
		Int64: *value,
		Valid: true,
	}
	return nil
}
func (t *NullInt64) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Int64)
}
