package service

import (
	"context"
	"go-rest-api/model/dto/request"
	"go-rest-api/model/dto/response"
)

type CategoryService interface {
	InsertCategory(ctx context.Context, req request.InsertCategoryReq) response.CategoryRes
	SearchCategoryById(ctx context.Context, categoryId int) response.CategoryRes
	SearchCategoryAll(ctx context.Context) []response.CategoryRes
	UpdateCategory(ctx context.Context, req request.UpdateCategoryReq) response.CategoryRes
	DeleteCategory(ctx context.Context, categoryId int)
}
