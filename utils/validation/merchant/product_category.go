package merchant

import "github.com/go-playground/validator/v10"

var productTypes = map[string]*struct{}{
	"Beverage": nil,
	"Food":     nil,
	"Snack":    nil,
	"Price":    nil,
	"ImageUrl": nil,
}

func validateProductCategoryImpl(s string) bool {
	_, ok := productTypes[s]
	return ok
}

func ValidateProductCategory(fl validator.FieldLevel) bool {
	return validateProductCategoryImpl(fl.Field().String())
}
