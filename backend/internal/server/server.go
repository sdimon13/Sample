package server

import (
	"context"
	"git.sample.ru/sample/internal/logger"
	"git.sample.ru/sample/internal/service"
	pb "git.sample.ru/sample/pkg/sample"
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Server struct {
	*pb.UnimplementedSampleServer
	*service.AppointmentService
	*service.ClientService
}

func (s *Server) Hello(ctx context.Context, in *pb.SampleRequest) (*pb.SampleResponse, error) {
	logger.Info.Print("Say Hello")

	return &pb.SampleResponse{
		Message: in.GetMessage(),
	}, nil
}

func (s *Server) GRPCRegister(gs grpc.ServiceRegistrar) {
	pb.RegisterSampleServer(gs, s)
}
func (s *Server) HTTPRegister(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterSampleHandler(ctx, mux, conn)
}
