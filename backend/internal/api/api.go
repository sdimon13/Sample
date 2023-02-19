package api

import (
	"context"
	"git.sample.ru/sample/internal/config"
	"git.sample.ru/sample/internal/logger"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"time"
)

type IServer interface {
	GRPCRegister(gs grpc.ServiceRegistrar)
	HTTPRegister(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type Api struct {
	c  *config.Config
	s  []IServer
	m  *runtime.ServeMux
	gs *grpc.Server
	hs *http.Server
}

func Get(c *config.Config, s []IServer, stream []grpc.StreamServerInterceptor, unary []grpc.UnaryServerInterceptor, m *runtime.ServeMux) *Api {
	return &Api{
		c: c,
		s: s,
		m: m,
		gs: grpc.NewServer(
			grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
				stream...,
			)),
			grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
				unary...,
			)),
		),
		hs: &http.Server{
			Addr:         c.HttpPort,
			ErrorLog:     logger.Error,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Handler:      cors(m),
		},
	}
}

func (a *Api) Run() error {
	lis, err := net.Listen("tcp", a.c.GrpcPort)
	if err != nil {
		logger.Error.Fatalln("Failed to listen:", err)
	}

	reflection.Register(a.gs)
	grpcPrometheus.Register(a.gs)
	logger.Info.Println("Serving gRPC on localhost" + a.c.GrpcPort)

	for _, server := range a.s {
		server.GRPCRegister(a.gs)
	}

	go func() {
		err := a.gs.Serve(lis)
		if err != nil {
			logger.Error.Fatal(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+a.c.GrpcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error.Fatalln("Failed to dial server:", err)
	}

	ctx := context.Background()
	for _, server := range a.s {
		err = server.HTTPRegister(ctx, a.m, conn)
		if err != nil {
			logger.Error.Fatalln("Failed to register gateway:", err)
		}
	}

	a.hs.Handler = cors(a.m)

	logger.Info.Printf("Serving gRPC-Gateway on http://localhost%s\n", a.c.GrpcPort)
	logger.Info.Printf("Serving Swagger-UI on http://localhost%s/swagger-ui/", a.c.HttpPort)
	go func() {
		err := a.hs.ListenAndServe()
		if err != nil {
			logger.Error.Fatal(err)
		}
	}()

	return err
}

func (a *Api) Close() error {
	a.gs.GracefulStop()
	return a.hs.Close()
}

/*func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")

		h.ServeHTTP(w, r)
	})
}*/

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight request
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}

		// Handle other requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
		h.ServeHTTP(w, r)
	})
}
