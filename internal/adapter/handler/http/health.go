package http

import (
	"api/internal/adapter/storage/mongodb"
	"net/http"
)

type HealthHandler struct {
	svc mongodb.Service
}

func NewHealthHandler(svc mongodb.Service) *HealthHandler {
	return &HealthHandler{
		svc,
	}
}

func (uh *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	handleSuccess(w, uh.svc.Health())
}
