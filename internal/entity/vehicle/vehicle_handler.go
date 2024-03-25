// internal/vehicle/handler.go
package vehicle

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler estrutura para injetar dependências no handler, como o serviço de veículos.
type Handler struct {
	Service Service
}

// NewHandler cria uma nova instância de Handler com as dependências necessárias.
func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// RegisterRoutes registra as rotas de veículos no roteador.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/vehicles", h.CreateVehicle).Methods("POST")
	router.HandleFunc("/vehicles/{id}", h.GetVehicle).Methods("GET")
	router.HandleFunc("/vehicles", h.ListVehicles).Methods("GET")
	router.HandleFunc("/vehicles/{id}", h.UpdateVehicle).Methods("PUT")
	router.HandleFunc("/vehicles/{id}", h.DeleteVehicle).Methods("DELETE")
}

// CreateVehicle é um manipulador HTTP para criar um novo veículo.
func (h *Handler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var v Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateVehicle(r.Context(), &v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}

// GetVehicle é um manipulador HTTP para buscar um veículo por ID.
func (h *Handler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	v, err := h.Service.GetVehicle(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(v)
}

// ListVehicles é um manipulador HTTP para listar todos os veículos.
func (h *Handler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.Service.ListVehicles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicles)
}

// UpdateVehicle é um manipulador HTTP para atualizar um veículo existente.
func (h *Handler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var v Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Adicionando o contexto e ID como argumentos
	if err := h.Service.UpdateVehicle(r.Context(), id, &v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(v)
}

// DeleteVehicle é um manipulador HTTP para deletar um veículo.
func (h *Handler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.Service.DeleteVehicle(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
