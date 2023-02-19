package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NullTime struct {
	Data   sql.NullTime
	Format string
}

func NewNullTime(data *time.Time) NullTime {
	if data == nil {
		return NullTime{
			// Default time format. Copy from format.go file (t Time) String() method
			Format: "2006-01-02 15:04:05.999999999 -0700 MST",
		}
	}
	return NullTime{
		Data: sql.NullTime{
			Time:  *data,
			Valid: true,
		},
		// Default time format. Copy from format.go file (t Time) String() method
		Format: "2006-01-02 15:04:05.999999999 -0700 MST",
	}
}

func (t NullTime) SetFormat(format string) NullTime {
	t.Format = format
	return t
}

// Methods for the user

func (t *NullTime) Valid() bool {
	return t.Data.Valid
}
func (t *NullTime) Get() time.Time {
	return t.Data.Time
}
func (t *NullTime) GetPtr() *time.Time {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullTime) Reset() {
	t.Data = sql.NullTime{
		Time:  time.Unix(0, 0),
		Valid: false,
	}
}
func (t *NullTime) Set(data time.Time) {
	t.Data = sql.NullTime{
		Time:  data,
		Valid: true,
	}
}
func (t *NullTime) ToString(format string) string {
	if !t.Valid() {
		return ""
	}
	return t.Get().Format(format)
}
func (t *NullTime) ToStringPtr(format string) *string {
	if !t.Valid() {
		return nil
	}
	val := t.Get().Format(format)
	return &val
}

// Methods for use with proto structures

func NewNullTimeFromTimestamppb(data *timestamppb.Timestamp) NullTime {
	if data == nil || !data.IsValid() {
		return NewNullTime(nil)
	}
	date := data.AsTime()
	return NewNullTime(&date)
}
func (t *NullTime) GetTimestamppb() *timestamppb.Timestamp {
	if !t.Valid() {
		return nil
	}
	return timestamppb.New(t.Get())
}

// fmt

func (t *NullTime) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Time.Format(t.Format))
}

// SQL

func (t *NullTime) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *NullTime) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullTime) UnmarshalJSON(data []byte) error {
	var value *time.Time
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullTime{
			Time:  time.Unix(0, 0),
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullTime{
		Time:  *value,
		Valid: true,
	}
	return nil
}
func (t *NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Time.Format(t.Format))
}
