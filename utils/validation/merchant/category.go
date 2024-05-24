package merchant

import "github.com/go-playground/validator/v10"

var merchantTypes = map[string]*struct{}{
	"SmallRestaurant":       nil,
	"MediumRestaurant":      nil,
	"LargeRestaurant":       nil,
	"MerchandiseRestaurant": nil,
	"BoothKiosk":            nil,
	"ConvenienceStore":      nil,
}

func validateCategoryImpl(s string) bool {
	_, ok := merchantTypes[s]
	return ok
}

func ValidateCategory(fl validator.FieldLevel) bool {
	return validateCategoryImpl(fl.Field().String())
}
