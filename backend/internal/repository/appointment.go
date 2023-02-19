package repository

import (
	"context"
	"git.sample.ru/sample/internal/db"
	"git.sample.ru/sample/internal/entity"
)

type Appointment struct {
	db *db.DB
}

type Interface interface {
	Get(id int) (*entity.Appointment, error)
	List() (*[]entity.Appointment, error)
	Add(t *entity.Appointment) (*entity.Appointment, error)
	Delete(id int) (bool, error)
}

func NewAppointment(d *db.DB) *Appointment {
	return &Appointment{db: d}
}

func (r *Appointment) Get(id int) (*entity.Appointment, error) {
	e := entity.Appointment{}

	//err := r.db.Client.
	//	QueryRow(context.Background(), "SELECT id FROM test WHERE id = $1 AND deleted_at IS NULL", id).
	//	Scan(&e.Id)

	//if err != nil {
	//	return nil, err
	//}

	return &e, nil
}

func (r *Appointment) List() (*[]entity.Appointment, error) {
	var d []entity.Appointment

	rows, err := r.db.Client.Query(context.Background(), "SELECT id, master_id, status_id, appointment_date, appointment_time FROM samples.appointments WHERE deleted_at IS NULL and status_id = 1 order by id")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.Appointment
		err := rows.Scan(&e.Id, &e.MasterId, &e.StatusId, &e.AppointmentDate, &e.AppointmentTime)

		if err != nil {
			return nil, err
		}

		d = append(d, e)
	}

	return &d, nil
}

func (r *Appointment) Add(e *entity.Appointment) (*entity.Appointment, error) {
	//err := r.db.Client.
	//	QueryRow(context.Background(), "INSERT INTO test(status, created_at, updated_at) VALUES ($1, now(), now()) RETURNING id", e.Status).
	//	Scan(&e.Id)

	//if err != nil {
	//	return nil, err
	//}

	return e, nil
}

func (r *Appointment) Delete(id int) (bool, error) {
	if _, err := r.Get(id); err != nil {
		return false, err
	}

	//_, err := r.db.Client.
	//	Query(context.Background(), "UPDATE test SET deleted_at = now() WHERE id = $1", id)

	//if err != nil {
	//	return false, err
	//}

	return true, nil
}
