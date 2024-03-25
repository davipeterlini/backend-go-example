// internal/vehicle/service.go
package vehicle

import (
	"context"
	"errors"
)

// Service define a interface para o serviço de veículos.
type Service interface {
	CreateVehicle(ctx context.Context, v *Vehicle) error
	GetVehicle(ctx context.Context, id string) (*Vehicle, error)
	ListVehicles(ctx context.Context) ([]*Vehicle, error)
	UpdateVehicle(ctx context.Context, id string, v *Vehicle) error
	DeleteVehicle(ctx context.Context, id string) error
}

// VehicleService é a implementação concreta de Service.
type VehicleService struct {
	repo *VehicleRepository
}

// NewVehicleService cria uma nova instância de VehicleService.
func NewVehicleService(repo *VehicleRepository) *VehicleService {
	return &VehicleService{
		repo: repo,
	}
}

// CreateVehicle cria um novo veículo.
func (s *VehicleService) CreateVehicle(ctx context.Context, v *Vehicle) error {
	// Aqui você pode adicionar validações ou lógicas de negócios antes de criar o veículo
	if v.Name == "" || v.Model == "" {
		return errors.New("name and model are required fields")
	}

	// Chamando o repositório para criar um novo veículo
	return s.repo.Create(ctx, v)
}

// GetVehicle busca um veículo pelo ID.
func (s *VehicleService) GetVehicle(ctx context.Context, id string) (*Vehicle, error) {
	if id == "" {
		return nil, errors.New("the vehicle id is required")
	}

	// Chamando o repositório para buscar o veículo
	return s.repo.Read(ctx, id)
}

// ListVehicles retorna uma lista de todos os veículos.
func (s *VehicleService) ListVehicles(ctx context.Context) ([]*Vehicle, error) {
	// Chamando o repositório para obter a lista de veículos
	vehicles, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (s *VehicleService) UpdateVehicle(ctx context.Context, id string, v *Vehicle) error {
	// Verifique se o veículo existe antes da atualização
	existingVehicle, err := s.repo.Read(ctx, id)
	if err != nil {
		return err // Pode ser um erro de conexão ao banco de dados ou similar
	}
	if existingVehicle == nil {
		return nil
	}

	// Prossiga com a atualização, pois o veículo existe
	return s.repo.Update(ctx, id, v)
}

func (s *VehicleService) DeleteVehicle(ctx context.Context, id string) error {
	// Verifique se o veículo existe antes da deleção
	existingVehicle, err := s.repo.Read(ctx, id)
	if err != nil {
		return err // Pode ser um erro de conexão ao banco de dados ou similar
	}
	if existingVehicle == nil {
		return nil
	}

	// Prossiga com a deleção, pois o veículo existe
	return s.repo.Delete(ctx, id)
}
