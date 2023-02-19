package router

import (
	"git.sample.ru/sample/internal/handler"
	"git.sample.ru/sample/internal/logger"
	"git.sample.ru/sample/pkg/swagger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type Router struct {
	hh *handler.Healthcheck
	sh *swagger.Swagger
}

func NewRouter(hh *handler.Healthcheck, sh *swagger.Swagger) *Router {
	return &Router{
		hh: hh,
		sh: sh,
	}
}

func (ro *Router) Get(env string) *runtime.ServeMux {
	router := runtime.NewServeMux()

	err := router.HandlePath(http.MethodGet, "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.hh.HealthcheckHandler(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	err = router.HandlePath(http.MethodGet, "/swagger.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.sh.GetSwaggerJsonHandler(w, r, "backend/api/api.swagger.json")
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	err = router.HandlePath(http.MethodGet, "/services.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.sh.GetSwaggerJsonHandler(w, r, "backend/api/services.json")
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	err = router.HandlePath(http.MethodGet, "/calendar.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.sh.GetSwaggerJsonHandler(w, r, "backend/api/calendar.json")
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	err = router.HandlePath(http.MethodGet, "/swagger-ui/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./backend/third_party/swagger-ui"))).ServeHTTP(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	return router
}
