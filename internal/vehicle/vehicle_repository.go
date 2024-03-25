package vehicle

import (
	"context"
	"sales/vehicle/internal/generics/repository"
)

type VehicleRepository struct {
	baseRepo *repository.BaseRepository
}

func NewVehicleRepository(baseRepo *repository.BaseRepository) *VehicleRepository {
	return &VehicleRepository{
		baseRepo: baseRepo,
	}
}

func (vr *VehicleRepository) Create(ctx context.Context, v *Vehicle) error {
	columns := []string{"name", "model", "status", "color", "mileage", "body_type", "transmission", "fuel_type", "doors", "review", "price", "description"}
	values := []interface{}{v.Name, v.Model, v.Status, v.Color, v.Mileage, v.BodyType, v.Transmission, v.FuelType, v.Doors, v.Review, v.Price, v.Description}
	return vr.baseRepo.Create(ctx, "vehicles", columns, values)
}

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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (vr *VehicleRepository) Update(ctx context.Context, id string, v *Vehicle) error {
	columns := []string{"name", "model", "status", "color", "mileage", "body_type", "transmission", "fuel_type", "doors", "review", "price", "description"}
	values := []interface{}{v.Name, v.Model, v.Status, v.Color, v.Mileage, v.BodyType, v.Transmission, v.FuelType, v.Doors, v.Review, v.Price, v.Description}
	return vr.baseRepo.Update(ctx, "vehicles", "id", id, columns, values)
}

func (vr *VehicleRepository) Delete(ctx context.Context, id string) error {
	return vr.baseRepo.Delete(ctx, "vehicles", "id", id)
}
