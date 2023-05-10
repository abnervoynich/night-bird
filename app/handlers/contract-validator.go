package handlers

import (
	"fmt"
	"github.com/abnervoynich/night-bird/app/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ValidatorHandler struct {
	ValidatorService *services.ContractValidator
}

func NewValidatorHandler(service *services.ContractValidator) *ValidatorHandler {
	return &ValidatorHandler{ValidatorService: service}
}

func (h *ValidatorHandler) GetRouter() http.Handler {
	r := chi.NewRouter()
	baseRoute := "/start"

	r.Post(baseRoute, h.Start)

	return r
}

func (h *ValidatorHandler) Start(w http.ResponseWriter, r *http.Request) {
	err := h.ValidatorService.StartAPIResponsesValidation()
	if err != nil {
		//TODO: add returning error message
		http.Error(w, http.StatusText(500), 500)
	}
	w.Write([]byte(fmt.Sprint("validation started..:%s")))
}
