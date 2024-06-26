package vehicle

import (
	"encoding/json"
	"net/http"
	"sales/vehicle/pkg/httperror"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) RegisterVehicleRoutes(router *mux.Router) {
	router.Handle("/vehicles", httperror.ErrorHandlerInterceptor(http.HandlerFunc(h.CreateVehicle))).Methods("POST")
	router.Handle("/vehicles/{id}", httperror.ErrorHandlerInterceptor(http.HandlerFunc(h.GetVehicle))).Methods("GET")
	router.Handle("/vehicles", httperror.ErrorHandlerInterceptor(http.HandlerFunc(h.ListVehicles))).Methods("GET")
	router.Handle("/vehicles/{id}", httperror.ErrorHandlerInterceptor(http.HandlerFunc(h.UpdateVehicle))).Methods("PUT")
	router.Handle("/vehicles/{id}", httperror.ErrorHandlerInterceptor(http.HandlerFunc(h.DeleteVehicle))).Methods("DELETE")
}

func (h *Handler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var v Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateVehicle(r.Context(), &v); err != nil {
		httperror.HandleCreateError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	v, err := h.Service.GetVehicle(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.Service.ListVehicles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicles)
}

func (h *Handler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var v Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateVehicle(r.Context(), id, &v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if err == nil {
		// TODO - FIX
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.Service.DeleteVehicle(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if err == nil {
		// TODO - FIX
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
