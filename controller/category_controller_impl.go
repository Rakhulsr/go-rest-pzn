package controller

import (
	"net/http"
	"strconv"

	"github.com/Rakhulsr/go-rest-api-pzn/helper"
	"github.com/Rakhulsr/go-rest-api-pzn/model/web"
	"github.com/Rakhulsr/go-rest-api-pzn/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateReq := web.CategoryCreateRequest{}
	helper.ParseJson(request, &categoryCreateReq)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteJson(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateReq := web.CategoryUpdateRequest{}

	helper.ParseJson(request, &categoryUpdateReq)

	categoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)

	categoryUpdateReq.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteJson(writer, webResponse)

}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteJson(writer, webResponse)

}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteJson(writer, webResponse)

}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteJson(writer, webResponse)

}
