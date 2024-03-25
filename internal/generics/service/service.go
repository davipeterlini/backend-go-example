/* // internal/vehicle/generic_service.go
package vehicle

import (
	"context"
	"errors"
)

// GenericService defines the interface for a generic service.
type GenericService interface {
	Create(ctx context.Context, entity interface{}) error
	Get(ctx context.Context, id string) (interface{}, error)
	List(ctx context.Context) ([]interface{}, error)
	Update(ctx context.Context, id string, entity interface{}) error
	Delete(ctx context.Context, id string) error
}

// genericService is the concrete implementation of GenericService.
type genericService struct {
	repo GenericRepository // Assume you have a GenericRepository interface similar to VehicleRepository
}

// NewGenericService creates a new instance of a generic service.
func NewGenericService(repo repository.GenericRepository) GenericService {
	return &genericService{
		repo: repo,
	}
}

// Create creates a new entity.
func (s *genericService) Create(ctx context.Context, entity interface{}) error {
	// Here you can add validations or business logic before creating the entity
	// Note: You will need to assert the type if specific fields need to be validated
	return s.repo.Create(ctx, entity)
}

// Get retrieves an entity by ID.
func (s *genericService) Get(ctx context.Context, id string) (interface{}, error) {
	if id == "" {
		return nil, errors.New("the entity id is required")
	}

	// Calling the repository to fetch the entity
	return s.repo.Read(ctx, id)
}

// List returns a list of all entities.
func (s *genericService) List(ctx context.Context) ([]interface{}, error) {
	// Calling the repository to get the list of entities
	entities, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// Update updates the details of an existing entity.
func (s *genericService) Update(ctx context.Context, id string, entity interface{}) error {
	// Here you can add validations or business logic before updating the entity
	// Note: You will need to assert the type if specific fields need to be validated
	return s.repo.Update(ctx, id, entity)
}

// Delete removes an entity from the database.
func (s *genericService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("the entity id is required for deletion")
	}

	// Calling the repository to remove the entity
	return s.repo.Delete(ctx, id)
} */