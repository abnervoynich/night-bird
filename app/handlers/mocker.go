package handlers

import (
	"encoding/json"
	"github.com/abnervoynich/night-bird/app/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type MockerHandler struct {
	MockerService *services.Mocker
}

func NewMockerHandler(mockerService *services.Mocker) *MockerHandler {
	return &MockerHandler{MockerService: mockerService}
}

func (h *MockerHandler) GetRouter() http.Handler {
	r := chi.NewRouter()
	baseRoute := "/serve"

	r.Get(baseRoute, h.ServeMockResponse) //the plan is to use something like "http://localhost/destination-url=https://linkedin.api.com/args
	r.Post(baseRoute, h.ServeMockResponse)
	r.Put(baseRoute, h.ServeMockResponse)
	r.Delete(baseRoute, h.ServeMockResponse)
	r.Patch(baseRoute, h.ServeMockResponse)

	return r
}

func (h *MockerHandler) ServeMockResponse(w http.ResponseWriter, r *http.Request) {
	response, err := h.MockerService.ServeMockResponse(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for headerKey, headerVal := range *response.Headers {
		w.Header().Add(headerKey, headerVal)
	}

	w.WriteHeader(response.ResponseCode)
	err = json.NewEncoder(w).Encode(response.Data)
	if err != nil {
		//TODO: add error returning message
	}
}
