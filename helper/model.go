package helper

import (
	"github.com/Rakhulsr/go-rest-api-pzn/model/domain"
	"github.com/Rakhulsr/go-rest-api-pzn/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}

	return categoriesResponse
}
