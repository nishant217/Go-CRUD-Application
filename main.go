package main

import (
	"CRUD/config"
	"CRUD/controller"
	"CRUD/helper"
	"CRUD/model"
	"CRUD/repository"
	"CRUD/router"
	"CRUD/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server")
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	tagRepository := repository.NewTagsRepositoryImpl(db)

	tagService := services.NewTagsServiceImpl(tagRepository, validate)
	tagController := controller.NewTagsController(tagService)

	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
