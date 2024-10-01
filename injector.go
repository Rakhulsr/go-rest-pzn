//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Rakhulsr/go-rest-api-pzn/app"
	"github.com/Rakhulsr/go-rest-api-pzn/controller"
	"github.com/Rakhulsr/go-rest-api-pzn/middleware"
	"github.com/Rakhulsr/go-rest-api-pzn/repositories"
	"github.com/Rakhulsr/go-rest-api-pzn/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repositories.NewCategoryRepository,
	wire.Bind(new(repositories.CategoryRepository), new(*repositories.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func provideValidator() *validator.Validate {

	return validator.New()
}

func InitServer() *http.Server {
	wire.Build(
		app.NewDb,
		provideValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewMiddleware,
		NewServer,
	)

	return nil
}
