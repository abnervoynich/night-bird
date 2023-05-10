package handlers

import (
	"fmt"
	"github.com/abnervoynich/night-bird/app/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AdminHandler struct {
	AdminService *services.AdminService
}

func NewAdminHandler(service *services.AdminService) *AdminHandler {
	return &AdminHandler{AdminService: service}
}

func (h *AdminHandler) GetRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.ServeAdminIndex)
	r.Get("/pellets", h.ServePelletsIndex)

	return r
}

// AdminOnly serves ui
func (h *AdminHandler) ServeAdminIndex(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}

// AdminOnly serves ui
func (h *AdminHandler) ServePelletsIndex(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}

// SavePellet saves pellet (for request or for response)
func (h *AdminHandler) SavePellet(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}

// DeletePellet saves pellet
func (h *AdminHandler) DeletePellet(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}

// GetPellet gets single pellet (for request or for response)
func (h *AdminHandler) GetPellet(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}

// GetPellets gets all pellets (for request or for response)
func (h *AdminHandler) GetPellets(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}

// TriggerValidationContract triggers validation contracts using response pellets
func (h *AdminHandler) TriggerValidationContract(w http.ResponseWriter, r *http.Request) {

	//TODO: implement calls for admin UI

	w.Write([]byte(fmt.Sprint("not implemented")))

}
