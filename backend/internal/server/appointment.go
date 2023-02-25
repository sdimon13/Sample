package server

import (
	"context"
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/logger"
	"git.sample.ru/sample/pkg/golibs/nulltype"
	pb "git.sample.ru/sample/pkg/sample"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
)

func (s *Server) AppointmentList(ctx context.Context, in *pb.AppointmentServiceListRequest) (*pb.AppointmentServiceListResponse, error) {
	ul, err := s.AppointmentService.GetAppointmentList(ctx, 1)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Appointments not found")
	}

	// Создаем словарь для хранения доступных времен по датам
	datesMap := make(map[string][]*pb.TimeSlot)
	// Создаем массив для хранения доступных дат
	dates := make([]string, 0)
	// Проходим по массиву структур Appointment и заполняем словарь доступных времен по датам
	for _, appointment := range ul {
		// Форматируем дату в формат "YYYY-MM-DD"
		date := appointment.AppointmentDate.Format("2006-01-02")
		// Форматируем время в формат "HH:mm"
		time := appointment.AppointmentTime.Format("15:04")
		// Если для текущей даты еще не были добавлены доступные времена, то создаем новую запись в словаре и добавляем
		// дату в массив доступных дат
		if _, ok := datesMap[date]; !ok {
			datesMap[date] = make([]*pb.TimeSlot, 0)
			dates = append(dates, date)
		}
		// Добавляем доступное время для текущей даты в словарь
		datesMap[date] = append(datesMap[date], &pb.TimeSlot{
			Id:   appointment.Id,
			Time: time,
		})
	}
	// Создаем словарь для хранения доступных времен по датам в формате, который соответствует протобуферной структуре
	// AvailableTime
	availableTimesMap := make(map[string]*pb.AvailableTime)
	// Проходим по словарю доступных времен по датам и конвертируем его в соответствующий протобуферный формат
	for date, times := range datesMap {
		availableTimesMap[date] = &pb.AvailableTime{TimeSlots: times}
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

func (s *Server) AdminAppointmentList(ctx context.Context, in *pb.AppointmentServiceListRequest) (*pb.AdminAppointmentServiceListResponse, error) {
	ul, err := s.AppointmentService.GetAppointmentList(ctx, 2)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Appointments not found")
	}

	// Создаем словарь для хранения доступных времен по датам
	datesMap := make(map[string][]*pb.AdminTimeSlot)
	// Создаем массив для хранения доступных дат
	dates := make([]string, 0)
	// Проходим по массиву структур Appointment и заполняем словарь доступных времен по датам
	for _, appointment := range ul {
		client, _ := s.ClientService.GetClient(ctx, appointment.ClientId.Get())

		// Форматируем дату в формат "YYYY-MM-DD"
		date := appointment.AppointmentDate.Format("2006-01-02")
		// Форматируем время в формат "HH:mm"
		time := appointment.AppointmentTime.Format("15:04")
		// Если для текущей даты еще не были добавлены доступные времена, то создаем новую запись в словаре и добавляем
		// дату в массив доступных дат
		if _, ok := datesMap[date]; !ok {
			datesMap[date] = make([]*pb.AdminTimeSlot, 0)
			dates = append(dates, date)
		}
		// Добавляем доступное время для текущей даты в словарь
		datesMap[date] = append(datesMap[date], &pb.AdminTimeSlot{
			Id:    appointment.Id,
			Time:  time,
			Name:  client.Name,
			Phone: client.Phone,
		})
	}
	// Создаем словарь для хранения доступных времен по датам в формате, который соответствует протобуферной структуре
	// AvailableTime
	availableTimesMap := make(map[string]*pb.AdminAvailableTime)
	// Проходим по словарю доступных времен по датам и конвертируем его в соответствующий протобуферный формат
	for date, times := range datesMap {
		availableTimesMap[date] = &pb.AdminAvailableTime{TimeSlots: times}
	}
	// Сортируем массив доступных дат в порядке возрастания
	sort.Strings(dates)
	// Создаем новую протобуферную структуру AppointmentListResponse с заполненными полями и возвращаем ее
	response := &pb.AdminAppointmentServiceListResponse{
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

func (s *Server) AppointmentUpdate(ctx context.Context, in *pb.AppointmentServiceUpdateRequest) (*pb.AppointmentServiceGetResponse, error) {
	logger.Info.Println("in")
	logger.Info.Println(in)
	logger.Info.Printf("id: %d, date: %s, time: %s, name: %s, phoneNumber: %s", in.Id, in.Date, in.Time, in.Name, in.PhoneNumber)

	ea, err := s.AppointmentService.GetAppointment(ctx, in.Id)
	if err != nil || ea.StatusId != 1 {
		return nil, status.Error(codes.NotFound, "Appointment not found")
	}

	ec, err := s.ClientService.FindClient(ctx, in.PhoneNumber, in.Name)
	if err != nil {
		c := entity.Client{
			Phone: in.PhoneNumber,
			Name:  in.Name,
		}

		ec, err = s.ClientService.AddClient(ctx, &c)
		if err != nil {
			return nil, status.Error(codes.NotFound, "Client cant create")
		}
	}

	ea.ClientId = nulltype.NewNullInt64(&ec.Id)
	ea.Note = nulltype.NewNullString(&in.Note)
	ea.StatusId = 2

	ea, err = s.AppointmentService.UpdateAppointment(ctx, ea)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Appointment cant update")
	}

	return &pb.AppointmentServiceGetResponse{
		Id:          ea.Id,
		Date:        ea.AppointmentDate.Format("2006-01-02"),
		Time:        ea.AppointmentTime.Format("15:04"),
		StatusId:    ea.StatusId,
		Name:        ec.Name,
		PhoneNumber: ec.Phone,
	}, nil
}
