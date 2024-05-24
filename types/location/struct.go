package location

type Location struct {
	Latitude  *float64 `json:"lat"  validate:"required"`
	Longitude *float64 `json:"long" validate:"required"`
}
