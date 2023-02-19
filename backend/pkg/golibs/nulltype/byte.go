package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullByte struct {
	Data sql.NullByte
}

func NewNullByte(data *byte) NullByte {
	if data == nil {
		return NullByte{}
	}
	return NullByte{
		Data: sql.NullByte{
			Byte:  *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *NullByte) Valid() bool {
	return t.Data.Valid
}
func (t *NullByte) Get() byte {
	return t.Data.Byte
}
func (t *NullByte) GetPtr() *byte {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullByte) Reset() {
	t.Data = sql.NullByte{
		Byte:  0,
		Valid: false,
	}
}
func (t *NullByte) Set(data byte) {
	t.Data = sql.NullByte{
		Byte:  data,
		Valid: true,
	}
}

// fmt

func (t *NullByte) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Byte)
}

// SQL

func (t *NullByte) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullByte) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullByte) UnmarshalJSON(data []byte) error {
	var value *byte
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullByte{
			Byte:  0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullByte{
		Byte:  *value,
		Valid: true,
	}
	return nil
}
func (t *NullByte) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Byte)
}
