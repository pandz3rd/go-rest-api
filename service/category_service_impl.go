package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/model/dao"
	"go-rest-api/model/dto/request"
	"go-rest-api/model/dto/response"
	"go-rest-api/repository"
	"go-rest-api/util"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validator          *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validator *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		validator:          validator,
	}
}

func (service *CategoryServiceImpl) InsertCategory(ctx context.Context, req request.InsertCategoryReq) response.CategoryRes {
	err := service.validator.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	category := dao.Category{
		Name: req.Name,
	}

	category = service.CategoryRepository.Insert(ctx, tx, category)

	return util.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) SearchCategoryById(ctx context.Context, categoryId int) response.CategoryRes {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	categoryExist, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}
	return util.ToCategoryResponse(categoryExist)
}

func (service *CategoryServiceImpl) SearchCategoryAll(ctx context.Context) []response.CategoryRes {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	categoryList := service.CategoryRepository.FindAll(ctx, tx)
	return util.ToListCategoryResponse(categoryList)
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, req request.UpdateCategoryReq) response.CategoryRes {
	err := service.validator.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	// If exist
	category, err := service.CategoryRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	category.Name = req.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return util.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category.Id)
}
