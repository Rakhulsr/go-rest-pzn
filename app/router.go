package app

import (
	"github.com/Rakhulsr/go-rest-api-pzn/controller"
	"github.com/Rakhulsr/go-rest-api-pzn/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(category controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", category.FindAll)
	router.GET("/api/categories/:categoryId", category.FindById)
	router.POST("/api/categories", category.Create)
	router.PUT("/api/categories/:categoryId", category.Update)
	router.DELETE("/api/categories/:categoryId", category.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
