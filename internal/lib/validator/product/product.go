package product

import (
	"errors"
	"reflect"

	"github.com/Marif226/melon/internal/model"
	"github.com/go-playground/validator/v10"
)

func Create(product *model.Product) error {
	validate := validator.New()
	err := validate.Struct(product)
	if err != nil {
		return err
	}

	return nil
}

func ValidateID(id int) error {
	if id < 0 {
		return errors.New("invalid ID: cannot be less or equal to 0")
	}

	return nil
}

func Update(request model.ProductUpdateRequest) error {
	v := reflect.ValueOf(request)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		switch fieldValue.Kind() {
		case reflect.String:
			if fieldValue.String() != "" {
				return nil
			}
		case reflect.Int:
			if fieldValue.Int() != 0 {
				return nil
			}
		case reflect.Float32:
			if fieldValue.Int() != 0.0 {
				return nil
			}
		}
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return err
	}

	return errors.New("at least one field must be non-empty")
}