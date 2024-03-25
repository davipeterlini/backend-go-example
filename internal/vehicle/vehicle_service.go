// internal/vehicle/service.go
package vehicle

import (
	"context"
	"errors"
)

// TODO - make excepetion errors for this class

type Service interface {
	CreateVehicle(ctx context.Context, v *Vehicle) error
	GetVehicle(ctx context.Context, id string) (*Vehicle, error)
	ListVehicles(ctx context.Context) ([]*Vehicle, error)
	UpdateVehicle(ctx context.Context, id string, v *Vehicle) error
	DeleteVehicle(ctx context.Context, id string) error
}

type VehicleService struct {
	repo *VehicleRepository
}

func NewVehicleService(repo *VehicleRepository) *VehicleService {
	return &VehicleService{
		repo: repo,
	}
}

func (s *VehicleService) CreateVehicle(ctx context.Context, v *Vehicle) error {
	if v.Name == "" || v.Model == "" {
		return errors.New("name and model are required fields")
	}

	return s.repo.Create(ctx, v)
}

func (s *VehicleService) GetVehicle(ctx context.Context, id string) (*Vehicle, error) {
	if id == "" {
		return nil, errors.New("the vehicle id is required")
	}

	return s.repo.Read(ctx, id)
}

func (s *VehicleService) ListVehicles(ctx context.Context) ([]*Vehicle, error) {
	vehicles, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (s *VehicleService) UpdateVehicle(ctx context.Context, id string, v *Vehicle) error {
	existingVehicle, err := s.repo.Read(ctx, id)
	if err != nil {
		return err
	}
	if existingVehicle == nil {
		return nil
	}
	return s.repo.Update(ctx, id, v)
}

func (s *VehicleService) DeleteVehicle(ctx context.Context, id string) error {
	existingVehicle, err := s.repo.Read(ctx, id)
	if err != nil {
		return err
	}
	if existingVehicle == nil {
		return nil
	}
	return s.repo.Delete(ctx, id)
}
