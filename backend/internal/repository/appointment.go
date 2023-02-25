package repository

import (
	"context"
	"git.sample.ru/sample/internal/db"
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/logger"
	qb "git.sample.ru/sample/pkg/golibs/query-builder"
	"github.com/georgysavva/scany/pgxscan"
	"sync"
)

type Appointment struct {
	db *db.DB
}

type AppointmentInterface interface {
	Get(ctx context.Context, id int64) (*entity.Appointment, error)
	List(ctx context.Context, statusId int32) ([]*entity.Appointment, error)
	Add(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error)
	Update(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

func NewAppointment(d *db.DB) *Appointment {
	return &Appointment{db: d}
}

func (r *Appointment) Get(ctx context.Context, id int64) (*entity.Appointment, error) {
	eu := entity.Appointment{}

	q := qb.NewQB().
		Columns("id", "master_id", "client_id", "service_id", "status_id", "appointment_date", "appointment_time", "note").
		From(eu.GetTable())
	q.Where().AddExpression("deleted_at IS NULL").
		AddExpression("id = ?", id)

	err := pgxscan.Get(ctx, r.db.Client, &eu, q.String(), q.GetArguments()...)
	if err != nil {
		return nil, err
	}

	return &eu, nil
}

func (r *Appointment) List(ctx context.Context, statusId int32) ([]*entity.Appointment, error) {
	var el []*entity.Appointment
	t := entity.Appointment{}.GetTable()

	q := qb.NewQB().
		Columns("id", "master_id", "client_id", "service_id", "status_id", "appointment_date", "appointment_time", "note").
		From(t)
	q.Where().
		AddExpression("status_id = ?", statusId).
		AddExpression("deleted_at IS NULL")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := pgxscan.Select(ctx, r.db.Client, &el, q.String(), q.GetArguments()...)
		if err != nil {
			logger.Error.Println("error on Appointment::get", err)
		}
	}()

	wg.Wait()

	return el, nil
}

func (r *Appointment) Add(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error) {
	err := r.db.Client.
		QueryRow(ctx, "INSERT INTO samples.appointments(master_id, client_id, service_id, status_id, appointment_date, appointment_time, note, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now()) RETURNING id", e.MasterId, e.ClientId, e.ServiceId, e.StatusId, e.AppointmentDate, e.AppointmentTime, e.Note).
		Scan(&e.Id)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *Appointment) Update(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error) {
	if _, err := r.Get(ctx, e.Id); err != nil {
		return nil, err
	}

	_, err := r.db.Client.
		Query(
			ctx, `UPDATE samples.appointments SET client_id = $2, status_id = $3, note = $4, updated_at = now() WHERE id = $1`, e.Id, e.ClientId.Get(), e.StatusId, e.Note.Get())

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *Appointment) Delete(ctx context.Context, id int64) (bool, error) {
	if _, err := r.Get(ctx, id); err != nil {
		return false, err
	}

	_, err := r.db.Client.
		Query(ctx, "UPDATE samples.appointments SET deleted_at = now() WHERE id = $1", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
