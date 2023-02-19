package server

import (
	"context"
	"git.sample.ru/sample/internal/logger"
	pb "git.sample.ru/sample/pkg/sample"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
)

func (s *Server) AppointmentList(ctx context.Context, in *pb.AppointmentServiceListRequest) (*pb.AppointmentServiceListResponse, error) {
	logger.Info.Print("Say Appointment list")

	ul, err := s.AppointmentService.GetAppointmentList()
	if err != nil {
		logger.Error.Printf("Error get appointment list: %s", err)
		return nil, status.Error(codes.NotFound, "Appointments not found")
	}

	// Создаем словарь для хранения доступных времен по датам
	datesMap := make(map[string][]string)
	// Создаем массив для хранения доступных дат
	dates := make([]string, 0)
	// Проходим по массиву структур Appointment и заполняем словарь доступных времен по датам
	for _, appointment := range *ul {
		// Форматируем дату в формат "YYYY-MM-DD"
		date := appointment.AppointmentDate.Format("2006-01-02")
		// Форматируем время в формат "HH:mm"
		time := appointment.AppointmentTime.Format("15:04")
		// Если для текущей даты еще не были добавлены доступные времена, то создаем новую запись в словаре и добавляем
		// дату в массив доступных дат
		if _, ok := datesMap[date]; !ok {
			datesMap[date] = make([]string, 0)
			dates = append(dates, date)
		}
		// Добавляем доступное время для текущей даты в словарь
		datesMap[date] = append(datesMap[date], time)
	}
	// Создаем словарь для хранения доступных времен по датам в формате, который соответствует протобуферной структуре
	// AvailableTime
	availableTimesMap := make(map[string]*pb.AvailableTime)
	// Проходим по словарю доступных времен по датам и конвертируем его в соответствующий протобуферный формат
	for date, times := range datesMap {
		availableTimesMap[date] = &pb.AvailableTime{Times: times}
	}
	// Сортируем массив доступных дат в порядке возрастания
	sort.Strings(dates)
	// Создаем новую протобуферную структуру AppointmentListResponse с заполненными полями и возвращаем ее
	response := &pb.AppointmentServiceListResponse{
		AvailableDates: dates,
		AvailableTimes: availableTimesMap,
	}

	return response, err
}

func (s *Server) AppointmentCreate(ctx context.Context, in *pb.AppointmentServiceCreateRequest) (*pb.AppointmentServiceGetResponse, error) {
	return &pb.AppointmentServiceGetResponse{
		Id:          0,
		Date:        in.Date,
		Time:        in.Time,
		StatusId:    2,
		Name:        in.Name,
		PhoneNumber: in.PhoneNumber,
	}, nil
}
