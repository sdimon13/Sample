package service

import (
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/repository"
)

type AppointmentService struct {
	r *repository.Appointment
}

type AppointmentInterface interface {
	GetTest(int) (*entity.Appointment, error)
	GetTestList() (*[]entity.Appointment, error)
	AddTest(message *entity.Appointment) (*entity.Test, error)
	DeleteTest(int) (bool, error)
}

func NewAppointmentService(r *repository.Appointment) *AppointmentService {
	return &AppointmentService{r}
}

func (s *AppointmentService) GetAppointment(id int) (*entity.Appointment, error) {
	return s.r.Get(id)
}

func (s *AppointmentService) GetAppointmentList() (*[]entity.Appointment, error) {
	return s.r.List()
}

func (s *AppointmentService) AddAppointment(m *entity.Appointment) (*entity.Appointment, error) {
	return s.r.Add(m)
}

func (s *AppointmentService) DeleteAppointment(id int) (bool, error) {
	return s.r.Delete(id)
}
