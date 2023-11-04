package product

import (
	"errors"

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