package vehicle

type Vehicle struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Model        string  `json:"model"`
	Status       string  `json:"status"` // Novo, Usado
	Color        string  `json:"color"`
	Mileage      int     `json:"mileage"`
	BodyType     string  `json:"bodyType"`
	Transmission string  `json:"transmission"` // Câmbio
	FuelType     string  `json:"fuelType"`
	Doors        int     `json:"doors"`
	Review       bool    `json:"review"`
	Price        float64 `json:"price"`
	Description  string  `json:"description"`
}
