package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NullDate struct {
	Data   sql.NullTime
	Format string
}

func NewNullDate(data *time.Time) NullDate {
	if data == nil {
		return NullDate{
			// Default time format. Copy from format.go file (t Time) String() method
			Format: "2006-01-02",
		}
	}
	return NullDate{
		Data: sql.NullTime{
			Time:  timeToDate(*data),
			Valid: true,
		},
		// Default time format. Copy from format.go file (t Time) String() method
		Format: "2006-01-02",
	}
}

func (t NullDate) SetFormat(format string) NullDate {
	t.Format = format
	return t
}

// Methods for the user

func (t *NullDate) Valid() bool {
	return t.Data.Valid
}
func (t *NullDate) Get() time.Time {
	return t.Data.Time
}
func (t *NullDate) GetPtr() *time.Time {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *NullDate) Reset() {
	t.Data = sql.NullTime{
		Time:  time.Unix(0, 0),
		Valid: false,
	}
}
func (t *NullDate) Set(data time.Time) {
	t.Data = sql.NullTime{
		Time:  timeToDate(data),
		Valid: true,
	}
}
func (t *NullDate) ToString(format string) string {
	if !t.Valid() {
		return ""
	}
	return t.Get().Format(format)
}
func (t *NullDate) ToStringPtr(format string) *string {
	if !t.Valid() {
		return nil
	}
	val := t.Get().Format(format)
	return &val
}

// Methods for use with proto structures

func NewNullDateFromTimestamppb(data *timestamppb.Timestamp) NullDate {
	if data == nil || !data.IsValid() {
		return NewNullDate(nil)
	}
	date := data.AsTime()
	return NewNullDate(&date)
}
func (t *NullDate) GetTimestamppb() *timestamppb.Timestamp {
	if !t.Valid() {
		return nil
	}
	return timestamppb.New(t.Get())
}

// fmt

func (t *NullDate) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Time.Format(t.Format))
}

// SQL

func (t *NullDate) Scan(value interface{}) error {
	err := t.Data.Scan(value)
	if err != nil {
		return err
	}
	if t.Valid() {
		t.Data.Time = timeToDate(t.Data.Time)
	}
	return nil
}
func (t *NullDate) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *NullDate) UnmarshalJSON(data []byte) error {
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
		Time:  timeToDate(*value),
		Valid: true,
	}
	return nil
}
func (t *NullDate) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Time.Format(t.Format))
}

// Utils

func timeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
