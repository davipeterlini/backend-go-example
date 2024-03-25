/* package hanlder

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

// GenericService defines a generic interface for our services
type GenericService interface {
	Create(ctx context.Context, entity interface{}) error
	Get(ctx context.Context, id string) (interface{}, error)
	List(ctx context.Context) ([]interface{}, error)
	Update(ctx context.Context, id string, entity interface{}) error
	Delete(ctx context.Context, id string) error
}

// Handler structure for injecting dependencies into the handler, such as a generic service.
type Handler struct {
	Service    GenericService
	EntityType reflect.Type // This is used to create new instances of the entity type
}

// NewHandler creates a new instance of Handler with the necessary dependencies.
func NewHandler(service GenericService, entityType reflect.Type) *Handler {
	return &Handler{
		Service:    service,
		EntityType: entityType,
	}
}

// Registers the routes for the entity on the router.
func (h *Handler) RegisterRoutes(router *mux.Router, path string) {
	router.HandleFunc(path, h.CreateEntity).Methods("POST")
	router.HandleFunc(path+"/{id}", h.GetEntity).Methods("GET")
	router.HandleFunc(path, h.ListEntities).Methods("GET")
	router.HandleFunc(path+"/{id}", h.UpdateEntity).Methods("PUT")
	router.HandleFunc(path+"/{id}", h.DeleteEntity).Methods("DELETE")
}

func (h *Handler) CreateEntity(w http.ResponseWriter, r *http.Request) {
	entity := reflect.New(h.EntityType).Interface() // Create a new entity of the specified type
	if err := json.NewDecoder(r.Body).Decode(entity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(r.Context(), entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity)
}

func (h *Handler) GetEntity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	entity, err := h.Service.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func (h *Handler) ListEntities(w http.ResponseWriter, r *http.Request) {
	entities, err := h.Service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entities)
}

func (h *Handler) UpdateEntity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	entity := reflect.New(h.EntityType).Elem().Interface() // Create a new instance of the entity type
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.Update(r.Context(), id, entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func (h *Handler) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.Service.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
*/