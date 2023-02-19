package server

import (
	"context"
	"git.sample.ru/sample/internal/logger"
	pb "git.sample.ru/sample/pkg/sample"
)

type Service struct {
	id    int32
	name  string
	price int32
	time  string
}

func (s *Server) ServiceList(ctx context.Context, in *pb.ServiceListRequest) (*pb.ServiceListResponse, error) {
	logger.Info.Print("Say Service list")

	var r []*pb.ServiceGetResponse

	services := &[]Service{
		{
			id:    1,
			name:  "Классический маникюр",
			price: 1000,
			time:  "60 минут",
		},
		{
			id:    2,
			name:  "Дизайн ногтей",
			price: 1500,
			time:  "90 минут",
		},
		{
			id:    3,
			name:  "Покрытие гель-лаком",
			price: 2000,
			time:  "120 минут",
		},
	}

	for _, s := range *services {
		r = append(r, mapServices(&s))
	}

	return &pb.ServiceListResponse{List: r}, nil
}

func mapServices(s *Service) *pb.ServiceGetResponse {
	return &pb.ServiceGetResponse{
		Id:    s.id,
		Name:  s.name,
		Price: s.price,
		Time:  s.time,
	}
}
