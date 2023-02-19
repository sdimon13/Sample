package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullBool struct {
	Data sql.NullBool
}

func NewNullBool(data *bool) NullBool {
	if data == nil {
		return NullBool{}
	}
	return NullBool{
		Data: sql.NullBool{
			Bool:  *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *NullBool) Valid() bool {
	return t.Data.Valid
}
func (t *NullBool) Get() bool {
	return t.Data.Bool
}
func (t *NullBool) GetPtr() *bool {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullBool) Reset() {
	t.Data = sql.NullBool{
		Bool:  false,
		Valid: false,
	}
}
func (t *NullBool) Set(data bool) {
	t.Data = sql.NullBool{
		Bool:  data,
		Valid: true,
	}
}

// fmt

func (t *NullBool) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Bool)
}

// SQL

func (t *NullBool) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullBool) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullBool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullBool{
			Bool:  false,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullBool{
		Bool:  *value,
		Valid: true,
	}
	return nil
}
func (t *NullBool) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Bool)
}
