package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"go_backend/config"
	"go_backend/controller"
	"go_backend/helper"
	"go_backend/model"
	"go_backend/repository"
	"go_backend/router"
	"go_backend/service"
	"net/http"
)

func main() {
	log.Info().Msg("Started Server!")
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("usr_rules").AutoMigrate(&model.UsrRules{})
	db.Table("process_rules").AutoMigrate(&model.ProcessRules{})

	userRulesRepository := repository.NewUsrRulesRepositoryImpl(db)
	processRulesRepository := repository.NewProcessRulesRepositoryImpl(db)

	userRulesService := service.NewUsrRulesServiceImpl(userRulesRepository, validate)
	processRulesService := service.NewProcessRulesServiceImpl(processRulesRepository, validate)

	userRulesController := controller.NewUsrRulesController(userRulesService)
	processRulesController := controller.NewProcessRulesController(processRulesService)

	routers := router.NewRouter(userRulesController, processRulesController)

	server := &http.Server{Addr: ":8888", Handler: routers}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
