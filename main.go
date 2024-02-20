package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/router"
	"golang-crud-gin/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})
	db.Table("user_details").AutoMigrate(&model.UserDetails{})

	// Repository
	usersRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	usersService := service.NewTagsServiceImpl(usersRepository, validate)

	// Controller
	usersController := controller.NewUsersController(usersService)

	// Router
	routes := router.NewRouter(usersController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
