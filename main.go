package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Rakhulsr/go-rest-api-pzn/app"
	"github.com/Rakhulsr/go-rest-api-pzn/controller"
	"github.com/Rakhulsr/go-rest-api-pzn/helper"
	"github.com/Rakhulsr/go-rest-api-pzn/middleware"
	"github.com/Rakhulsr/go-rest-api-pzn/repositories"
	"github.com/Rakhulsr/go-rest-api-pzn/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDb()
	validate := validator.New()

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewMiddleware(router),
	}
	log.Println("Listening to", "http://"+server.Addr)
	err := server.ListenAndServe()
	helper.PanicError(err)

}

func initDb(db *sql.DB) {
	err := db.Ping()
	helper.PanicError(err)

	log.Println("DB CONNECTED!")
}
