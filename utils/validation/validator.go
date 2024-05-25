package validation

import (
	"reflect"
	"strings"

	"github.com/JesseNicholas00/BeliMang/utils/validation/merchant"

	"github.com/JesseNicholas00/BeliMang/utils/validation/image"
	"github.com/JesseNicholas00/BeliMang/utils/validation/intlen"
	"github.com/JesseNicholas00/BeliMang/utils/validation/iso8601"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

var customFields = []customField{
	{
		Tag:       "imageExt",
		Validator: image.ValidateImageExtension,
	},
	{
		Tag:       "iso8601",
		Validator: iso8601.ValidateIso8601,
	},
	{
		Tag:       "intlen",
		Validator: intlen.ValidateIntLen,
	},
	{
		Tag:       "merchantCategory",
		Validator: merchant.ValidateCategory,
	},
	{
		Tag:       "productCategory",
		Validator: merchant.ValidateProductCategory,
	},
}

type customField struct {
	Tag       string
	Validator validator.Func
}

func NewEchoValidator() echo.Validator {
	validator := validator.New(validator.WithRequiredStructEnabled())

	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	for _, customField := range customFields {
		validator.RegisterValidation(customField.Tag, customField.Validator)
	}

	return &EchoValidator{
		validator: validator,
	}
}
