package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullFloat64 struct {
	Data sql.NullFloat64
}

func NewNullFloat64(data *float64) NullFloat64 {
	if data == nil {
		return NullFloat64{}
	}
	return NullFloat64{
		Data: sql.NullFloat64{
			Float64: *data,
			Valid:   true,
		},
	}
}

// Methods for the user

func (t *NullFloat64) Valid() bool {
	return t.Data.Valid
}
func (t *NullFloat64) Get() float64 {
	return t.Data.Float64
}
func (t *NullFloat64) GetPtr() *float64 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullFloat64) Reset() {
	t.Data = sql.NullFloat64{
		Float64: 0,
		Valid:   false,
	}
}
func (t *NullFloat64) Set(data float64) {
	t.Data = sql.NullFloat64{
		Float64: data,
		Valid:   true,
	}
}

// fmt

func (t *NullFloat64) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Float64)
}

// SQL

func (t *NullFloat64) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullFloat64) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullFloat64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullFloat64{
			Float64: 0,
			Valid:   false,
		}
		return nil
	}
	t.Data = sql.NullFloat64{
		Float64: *value,
		Valid:   true,
	}
	return nil
}
func (t *NullFloat64) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Float64)
}
