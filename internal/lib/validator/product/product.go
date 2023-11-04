package product

import (
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