package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
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
	logger             *logrus.Logger
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validator *validator.Validate, logger *logrus.Logger) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		validator:          validator,
		logger:             logger,
	}
}

func (service *CategoryServiceImpl) InsertCategory(ctx context.Context, req request.InsertCategoryReq) response.CategoryRes {
	traceId, _ := ctx.Value("traceId").(string)
	service.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start create new category: ", req.Name)
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
	traceId, _ := ctx.Value("traceId").(string)
	service.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start search category by id: ", categoryId)
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
	traceId, _ := ctx.Value("traceId").(string)
	service.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start search category all")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	categoryList := service.CategoryRepository.FindAll(ctx, tx)
	return util.ToListCategoryResponse(categoryList)
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, req request.UpdateCategoryReq) response.CategoryRes {
	traceId, _ := ctx.Value("traceId").(string)
	service.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start update new category: ", req.Id)
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
	traceId, _ := ctx.Value("traceId").(string)
	service.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start delete new category: ", categoryId)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.HandleRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category.Id)
}
