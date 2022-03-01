package router

import (
	"github.com/go-chi/chi/v5"
	"sample-api/internal/sms"
)

const (
	pathGetData = "/sms/data"
	pathAddData = "/sms/data"
)

func InitRouter(service sms.Service) *chi.Mux {
	r := chi.NewRouter()
	r.Get(pathGetData, service.GetData)
	r.Post(pathAddData, service.AddData)
	return r
}
