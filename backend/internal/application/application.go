package application

import (
	"context"
	"git.sample.ru/sample/internal/api"
	"git.sample.ru/sample/internal/config"
	"git.sample.ru/sample/internal/db"
	"git.sample.ru/sample/internal/handler"
	"git.sample.ru/sample/internal/logger"
	"git.sample.ru/sample/internal/repository"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"git.sample.ru/sample/internal/router"
	"git.sample.ru/sample/internal/server"
	"git.sample.ru/sample/internal/service"
	"git.sample.ru/sample/pkg/swagger"
	"google.golang.org/grpc"
	"time"
)

const (
	TypeApi     = "api"
	TypeCommand = "command"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
	Api *api.Api
	Ext Ext
}

type Ext struct {
	C *grpc.ClientConn
}

func Get(mode string) (*Application, error) {
	ctx := context.Background()
	cfg := config.Get()

	/* DB */
	dbc, err := db.Get(ctx, cfg.DSN)
	if err != nil {
		logger.Error.Fatal(err)
	}

	app := &Application{
		DB:  dbc,
		Cfg: cfg,
	}

	if err != nil {
		logger.Error.Fatal(err)
	}

	switch mode {
	case TypeApi:
		var servers []api.IServer
		var streamInterceptors []grpc.StreamServerInterceptor
		var unaryInterceptors []grpc.UnaryServerInterceptor

		/* Interceptors (middleware) setup */
		streamInterceptors = append(streamInterceptors, grpcPrometheus.StreamServerInterceptor)
		unaryInterceptors = append(unaryInterceptors, grpcPrometheus.UnaryServerInterceptor)

		hh := handler.GetHealthHandler(app.DB)
		hh.SetApplicationState("1.0.0", time.Now().Format("2006-01-02 15:04:05"))

		// Init repository
		a := repository.NewAppointment(dbc)

		// Init service
		as := service.NewAppointmentService(a)

		s := &server.Server{
			AppointmentService: as,
		}
		servers = append(servers, s)

		ro := router.NewRouter(hh, swagger.GetSwaggerHandler(cfg.SwaggerPath))
		app.Api = api.Get(cfg, servers, streamInterceptors, unaryInterceptors, ro.Get(cfg.AppEnv))

	case TypeCommand:
	}

	return app, nil
}

func (a *Application) Stop() error {
	//defer a.Ext.C.Close()

	return nil
}
