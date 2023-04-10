package handlers

import (
	"Car-parking-management-systems/models"
	"Car-parking-management-systems/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CarHandler struct {
	carService *services.CarService
}

func NewCarHandler(carService *services.CarService) *CarHandler {
	return &CarHandler{carService: carService}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Car models.Car
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.carService.CreateCar(&req.Car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"status": "success", "message": "Car created successfully!"}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respJSON)
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID := vars["id"]

	var req struct {
		ID  string `json:"id"`
		Car models.Car
	}

	req.ID = carID

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.carService.UpdateCar(req.ID, &req.Car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CarHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.carService.GetAllCars()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID := vars["id"]

	if err := h.carService.DeleteCar(carID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"status": "success", "message": "Car deleted successfully!"}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}
