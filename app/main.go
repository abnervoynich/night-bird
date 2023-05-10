package main

import (
	"fmt"
	"github.com/abnervoynich/night-bird/app/handlers"
	"github.com/abnervoynich/night-bird/app/models"
	"github.com/abnervoynich/night-bird/app/services"
	"github.com/abnervoynich/night-bird/app/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// loads pellets from local files
	localPellets, err := getPellets()
	if err != nil {
		log.Fatalf("cannot get local pellets: %s", err)
		return
	}

	// loads pellets flows from local files
	localPelletsFlows, err := getPelletsFlows()
	if err != nil {
		log.Fatalf("cannot get local pellets flows: %s", err)
		return
	}

	//initialize handler objects
	adminHandler := handlers.NewAdminHandler(services.NewAdminService())
	mockerHandler := handlers.NewMockerHandler(services.NewMockerService(localPellets))
	validatorHandler := handlers.NewValidatorHandler(services.NewContractValidatorService(localPellets, localPelletsFlows))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/admin", adminHandler.GetRouter())
	r.Mount("/mocker", mockerHandler.GetRouter())
	r.Mount("/validator", validatorHandler.GetRouter())
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error trying to start server:%s", err)
	}
}

// getPellets gets the local pellets and return a list of objects
func getPellets() ([]models.Pellet, error) {
	pelletsBytes, err := utils.GetLocalPellets()
	if err != nil {
		return nil, err
	}

	pellets, err := utils.ParseBytePellets(pelletsBytes)
	if err != nil {
		return nil, err
	}

	return pellets, nil
}

// getPelletsFlows gets the local pellets flows and return a list of objects
func getPelletsFlows() ([]models.PelletFlow, error) {
	pelletsFlowsBytes, err := utils.GetLocalPelletsFlows()
	if err != nil {
		return nil, err
	}

	pelletsFlows, err := utils.ParseBytePelletsFlows(pelletsFlowsBytes)
	if err != nil {
		return nil, err
	}

	return pelletsFlows, nil
}
