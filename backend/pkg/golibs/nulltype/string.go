package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullString struct {
	Data sql.NullString
}

func NewNullString(data *string) NullString {
	if data == nil {
		return NullString{}
	}
	return NullString{
		Data: sql.NullString{
			String: *data,
			Valid:  true,
		},
	}
}

// Methods for the user

func (t *NullString) Valid() bool {
	return t.Data.Valid
}
func (t *NullString) Get() string {
	return t.Data.String
}
func (t *NullString) GetPtr() *string {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullString) Reset() {
	t.Data = sql.NullString{
		String: "",
		Valid:  false,
	}
}
func (t *NullString) Set(data string) {
	t.Data = sql.NullString{
		String: data,
		Valid:  true,
	}
}

// fmt

func (t *NullString) String() string {
	if !t.Valid() {
		return ""
	}
	return t.Data.String
}

// SQL

func (t *NullString) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullString) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullString) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullString{
			String: "",
			Valid:  false,
		}
		return nil
	}
	t.Data = sql.NullString{
		String: *value,
		Valid:  true,
	}
	return nil
}
func (t *NullString) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.String)
}
