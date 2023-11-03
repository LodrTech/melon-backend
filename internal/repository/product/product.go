package product

import (
	"context"
	"errors"
	"fmt"
	"log"
	"github.com/Marif226/melon/internal/lib/querybuilder"
	"github.com/Marif226/melon/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type productRepo struct {
	db *pgx.Conn
}

func NewProductRepo(conn *pgx.Conn) *productRepo {
	return &productRepo {
		db: conn,
	}
}

func (r *productRepo) Create(ctx context.Context, request model.Product) (model.Product, error) {
	query, args, err := querybuilder.ProductCreate(request)
	if err != nil {
		return model.Product{}, nil
	}

	err = r.db.QueryRow(ctx, query, args...).Scan(&request.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			err = fmt.Errorf(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where))
			log.Println(err)
		}
		return model.Product{}, err
	}

	return request, nil
}