// internal/vehicle/vehicle_repository.go
package vehicle

import (
	"context"
	"sales/vehicle/internal/repository" // Substitua pelo caminho correto até o seu BaseRepository
)

// VehicleRepository oferece métodos específicos para trabalhar com veículos no banco de dados.
type VehicleRepository struct {
	baseRepo *repository.BaseRepository
}

// NewVehicleRepository cria uma nova instância de VehicleRepository.
func NewVehicleRepository(baseRepo *repository.BaseRepository) *VehicleRepository {
	return &VehicleRepository{
		baseRepo: baseRepo,
	}
}

// Create adiciona um novo veículo ao banco de dados.
func (vr *VehicleRepository) Create(ctx context.Context, v *Vehicle) error {
	columns := []string{"name", "model", "status", "color", "mileage", "body_type", "transmission", "fuel_type", "doors", "review", "price", "description"}
	values := []interface{}{v.Name, v.Model, v.Status, v.Color, v.Mileage, v.BodyType, v.Transmission, v.FuelType, v.Doors, v.Review, v.Price, v.Description}
	return vr.baseRepo.Create(ctx, "vehicles", columns, values)
}

// Read busca um veículo pelo ID.
func (vr *VehicleRepository) Read(ctx context.Context, id string) (*Vehicle, error) {
	row, err := vr.baseRepo.Read(ctx, "vehicles", "id", id)
	if err != nil {
		return nil, err
	}
	var v Vehicle
	if err := row.Scan(&v.ID, &v.Name, &v.Model, &v.Status, &v.Color, &v.Mileage, &v.BodyType, &v.Transmission, &v.FuelType, &v.Doors, &v.Review, &v.Price, &v.Description); err != nil {
		return nil, err
	}
	return &v, nil
}

// List retorna uma lista de todos os veículos do banco de dados.
func (vr *VehicleRepository) List(ctx context.Context) ([]*Vehicle, error) {
	columns := []string{"id", "name", "model", "status", "color", "mileage", "body_type", "transmission", "fuel_type", "doors", "review", "price", "description"}
	rows, err := vr.baseRepo.List(ctx, "vehicles", columns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []*Vehicle
	for rows.Next() {
		var v Vehicle
		if err := rows.Scan(&v.ID, &v.Name, &v.Model, &v.Status, &v.Color, &v.Mileage, &v.BodyType, &v.Transmission, &v.FuelType, &v.Doors, &v.Review, &v.Price, &v.Description); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &v)
	}

	// Verificar erros ao finalizar a iteração
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

// Update modifica um veículo existente.
func (vr *VehicleRepository) Update(ctx context.Context, id string, v *Vehicle) error {
	columns := []string{"name", "model", "status", "color", "mileage", "body_type", "transmission", "fuel_type", "doors", "review", "price", "description"}
	values := []interface{}{v.Name, v.Model, v.Status, v.Color, v.Mileage, v.BodyType, v.Transmission, v.FuelType, v.Doors, v.Review, v.Price, v.Description}
	return vr.baseRepo.Update(ctx, "vehicles", "id", id, columns, values)
}

// Delete remove um veículo do banco de dados.
func (vr *VehicleRepository) Delete(ctx context.Context, id string) error {
	return vr.baseRepo.Delete(ctx, "vehicles", "id", id)
}
