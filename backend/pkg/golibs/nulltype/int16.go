package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullInt16 struct {
	Data sql.NullInt16
}

func NewNullInt16(data *int16) NullInt16 {
	if data == nil {
		return NullInt16{}
	}
	return NullInt16{
		Data: sql.NullInt16{
			Int16: *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *NullInt16) Valid() bool {
	return t.Data.Valid
}
func (t *NullInt16) Get() int16 {
	return t.Data.Int16
}
func (t *NullInt16) GetPtr() *int16 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullInt16) Reset() {
	t.Data = sql.NullInt16{
		Int16: 0,
		Valid: false,
	}
}
func (t *NullInt16) Set(data int16) {
	t.Data = sql.NullInt16{
		Int16: data,
		Valid: true,
	}
}

// fmt

func (t *NullInt16) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Int16)
}

// SQL

func (t *NullInt16) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullInt16) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullInt16) UnmarshalJSON(data []byte) error {
	var value *int16
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullInt16{
			Int16: 0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullInt16{
		Int16: *value,
		Valid: true,
	}
	return nil
}
func (t *NullInt16) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Int16)
}
