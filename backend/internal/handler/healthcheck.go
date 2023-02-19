package handler

import (
	"context"
	"git.sample.ru/sample/internal/db"
	"git.sample.ru/sample/internal/logger"
	"git.sample.ru/sample/internal/utils"
	"net/http"
	"runtime"
)

type Healthcheck struct {
	db *db.DB
	as *State
}

type State struct {
	Version   string
	StartedAt string
}

type healthResponse struct {
	Status    string           `json:"status"`
	Db        dbStateResponse  `json:"db"`
	App       appStateResponse `json:"app"`
	RequestId string           `json:"request_id"`
}

type dbStateResponse struct {
	State         string `json:"state"`
	MaxConns      int32  `json:"max_conns"`
	AcquiredConns int32  `json:"acquired_conns"`
}

type appStateResponse struct {
	Version   string                 `json:"version"`
	StartedAt string                 `json:"started_at"`
	Memory    appStateMemoryResponse `json:"memory"`
}

type appStateMemoryResponse struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys        uint64 `json:"sys"`
	NumGC      uint32 `json:"num_gc"`
}

func GetHealthHandler(db *db.DB) *Healthcheck {
	return &Healthcheck{
		db,
		&State{},
	}
}

func (h Healthcheck) HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	response := &healthResponse{
		Status:    "OK",
		Db:        h.getDbState(),
		App:       h.getAppState(),
		RequestId: req.Header.Get("x-request-id"),
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *Healthcheck) getDbState() dbStateResponse {
	if h.db != nil {
		var dbStatus = "OK"
		err := h.db.Client.Ping(context.Background())
		dbStat := h.db.Client.Stat()
		if err != nil {
			logger.Error.Print(err)
			dbStatus = "FAIL"
		}

		return dbStateResponse{
			dbStatus,
			dbStat.MaxConns(),
			dbStat.AcquiredConns(),
		}
	}

	return dbStateResponse{
		State:         "FAIL",
		MaxConns:      0,
		AcquiredConns: 0,
	}
}

func (h *Healthcheck) getAppState() appStateResponse {
	var ms = runtime.MemStats{}
	runtime.ReadMemStats(&ms)

	return appStateResponse{
		Version:   h.as.Version,
		StartedAt: h.as.StartedAt,
		Memory: appStateMemoryResponse{
			Alloc:      bToMb(ms.Alloc),
			TotalAlloc: bToMb(ms.TotalAlloc),
			Sys:        bToMb(ms.Sys),
			NumGC:      ms.NumGC,
		},
	}
}

func (h *Healthcheck) SetApplicationState(version string, startedAt string) {
	h.as = &State{version, startedAt}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
