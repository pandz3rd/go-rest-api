package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/dao"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, category dao.Category) dao.Category {
	script := "insert into `category` (`name`) values (?)"
	result, err := tx.ExecContext(ctx, script, category.Name)
	helper.PanicIfError(err)
	id, _ := result.LastInsertId()
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (dao.Category, error) {
	script := "select id, name from `category` where `id` = ?"
	rows, err := tx.QueryContext(ctx, script, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := dao.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("Kaga ade")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []dao.Category {
	script := "select id, name from `category`"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	categoryList := []dao.Category{}
	for rows.Next() {
		category := dao.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categoryList = append(categoryList, category)
	}
	return categoryList
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category dao.Category) dao.Category {
	script := "update category set `name` = ? where `id` = ?"
	_, err := tx.ExecContext(ctx, script, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	script := "delete from category where `id` = ?"
	_, err := tx.ExecContext(ctx, script, categoryId)
	helper.PanicIfError(err)
}
