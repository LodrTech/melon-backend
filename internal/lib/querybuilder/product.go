package querybuilder

import (
	"github.com/Marif226/melon/internal/model"
	"github.com/Masterminds/squirrel"
)

func ProductCreate(request model.Product) (string, []any, error)  {
	qbuilder := squirrel.Insert(
		"products",
	).Columns(
		"name",
		"description",
		"price",
		"weight",
	).Values(
		request.Name,
		request.Description,
		request.Price,
		request.Weight,
	).Suffix(
		"RETURNING id",
	).PlaceholderFormat(squirrel.Dollar)

	return qbuilder.ToSql()
}