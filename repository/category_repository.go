package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/dao"
)

type CategoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, category dao.Category) dao.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (dao.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []dao.Category
	Update(ctx context.Context, tx *sql.Tx, category dao.Category) dao.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
}
