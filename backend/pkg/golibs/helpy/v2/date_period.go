package helpy

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/dimonrus/porterr"
	"net/http"
	"net/url"
	"time"
)

// Требуемый формат даты в фильтрах по периоду
const DateLayout = "2006-01-02"

type (
	// Тип временной границы периода
	Date time.Time
	// Период для выборки данных с часовым смещением(поясом)// Структура данных для функции UnmarshalJSON
	DatePeriod struct {
		// Начало
		Start *Date `json:"start"`
		// Окончание
		End *Date `json:"end"`
	}
)

func NewDatePeriod() *DatePeriod {
	return &DatePeriod{}
}

func NewDateFromTime(t *time.Time) *Date {
	if t == nil {
		return nil
	}
	d := Date(*t)
	return &d
}

// Установка верного времени для границ
func (p *DatePeriod) SetBoundaries() {
	if p.Start != nil {
		t := p.Start.ToTime()
		*t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
		*p.Start = Date(*t)
	}
	if p.End != nil {
		t := p.End.ToTime()
		*t = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999, time.Local)
		*p.End = Date(*t)
	}
	return
}

// Получение фильтра по дате из ГЕТ параметров
func (p *DatePeriod) ParseGetParams(filter url.Values) porterr.IError {
	e := porterr.New(porterr.PortErrorValidation, "Validation error").HTTP(http.StatusBadRequest)
	p.Start = parseDatePeriodValue(filter, "start", e)
	p.End = parseDatePeriodValue(filter, "end", e)
	e = e.IfDetails()
	if e != nil {
		return e
	}
	p.SetBoundaries()
	return nil
}

// Извлечение конкретной временной границы фильтра
func parseDatePeriodValue(filter url.Values, valueName string, e porterr.IError) *Date {
	keys, ok := filter[valueName]
	if ok && len(keys) == 1 {
		t, err := time.Parse(DateLayout, keys[0])
		if err != nil {
			e = e.PushDetail(porterr.PortErrorParam, valueName, err.Error())
		} else {
			d := Date(t)
			return &d
		}
	}
	return nil
}

// Анмаршал типа Date
func (d *Date) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	t, err := time.Parse(DateLayout, string(data[1:len(data)-1]))
	*d = Date(t)
	return err
}

// Маршалинг типа Date
func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToTime().Format(DateLayout))
}

// Конвертация Date в time.Time
func (d *Date) ToTime() *time.Time {
	if d != nil {
		t := time.Time(*d)
		return &t
	}
	return nil
}

// Для вставки значения в бд
func (d *Date) Value() (driver.Value, error) {
	t := d.ToTime()
	return *t, nil
}

// Для извлечения данных из бд
func (d *Date) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		*d = Date(t)
	}
	return nil
}
