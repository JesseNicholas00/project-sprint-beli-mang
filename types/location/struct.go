package location

type Location struct {
	Latitude  *float32 `json:"lat"  validate:"required"`
	Longitude *float32 `json:"long" validate:"required"`
}
