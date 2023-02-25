package service

import (
	"context"
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/repository"
)

type AppointmentService struct {
	r *repository.Appointment
}

type AppointmentInterface interface {
	GetAppointment(ctx context.Context, id int64) (*entity.Appointment, error)
	GetAppointmentList(ctx context.Context, statusId int32) ([]*entity.Appointment, error)
	AddAppointment(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error)
	UpdateAppointment(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error)
	DeleteAppointment(ctx context.Context, id int64) (bool, error)
}

func NewAppointmentService(r *repository.Appointment) *AppointmentService {
	return &AppointmentService{r}
}

func (s *AppointmentService) GetAppointment(ctx context.Context, id int64) (*entity.Appointment, error) {
	return s.r.Get(ctx, id)
}

func (s *AppointmentService) GetAppointmentList(ctx context.Context, statusId int32) ([]*entity.Appointment, error) {
	return s.r.List(ctx, statusId)
}

func (s *AppointmentService) AddAppointment(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error) {
	return s.r.Add(ctx, e)
}

func (s *AppointmentService) UpdateAppointment(ctx context.Context, e *entity.Appointment) (*entity.Appointment, error) {
	return s.r.Update(ctx, e)
}

func (s *AppointmentService) DeleteAppointment(ctx context.Context, id int64) (bool, error) {
	return s.r.Delete(ctx, id)
}
