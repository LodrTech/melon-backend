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

func ProductList(request model.ProductListRequest) (string, []any, error)  {
	qbuilder := squirrel.Select(
		"p.id",
		"p.name",
		"p.description",
		"p.price",
		"p.weight",
	).From(
		"products AS p",
	).PlaceholderFormat(squirrel.Dollar)

	return qbuilder.ToSql()
}

func ProductGet(id int) (string, []any, error)  {
	qbuilder := squirrel.Select(
		"p.id",
		"p.name",
		"p.description",
		"p.price",
		"p.weight",
	).From(
		"products AS p",
	).Where(
		squirrel.Eq{"p.id": id},
	).PlaceholderFormat(squirrel.Dollar)

	return qbuilder.ToSql()
}

func ProductUpdate(request model.ProductUpdateRequest) (string, []any, error)  {
	setMap := make(map[string]any)

	if request.Name != "" {
		setMap["name"] = request.Name
	}

	if request.Description != "" {
		setMap["description"] = request.Description
	}

	if request.Price != 0.0 {
		setMap["price"] = request.Price
	}

	if request.Weight != 0 {
		setMap["weight"] = request.Weight
	}

	qbuilder := squirrel.Update(
		"products AS p",
	).SetMap(
		setMap,
	).Where(
		squirrel.Eq{"p.id": request.ID},
	).Suffix(
		"RETURNING p.id, p.name, p.description, p.price, p.weight",
	).PlaceholderFormat(squirrel.Dollar)

	return qbuilder.ToSql()
}