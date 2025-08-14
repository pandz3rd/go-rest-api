package util

import (
	"go-rest-api/model/dao"
	"go-rest-api/model/dto/response"
)

func ToCategoryResponse(category dao.Category) response.CategoryRes {
	return response.CategoryRes{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToListCategoryResponse(categoryList []dao.Category) []response.CategoryRes {
	resultList := []response.CategoryRes{}
	for _, category := range categoryList {
		resultList = append(resultList, ToCategoryResponse(category))
	}
	return resultList
}
