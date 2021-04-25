package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-echo-postgresql-rest-api-example/config"
	"golang-echo-postgresql-rest-api-example/controller"
	_ "golang-echo-postgresql-rest-api-example/docs"
	"golang-echo-postgresql-rest-api-example/handler"
	"golang-echo-postgresql-rest-api-example/repository"
	"golang-echo-postgresql-rest-api-example/routes"
	"golang-echo-postgresql-rest-api-example/util"
	"log"
)

var userController *controller.UserController

// @title Golang User REST API
// @description Provides access to the core features of Golang User REST API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {
	e := echo.New()

	e.HTTPErrorHandler = handler.ErrorHandler
	e.Validator = util.NewValidationUtil()
	config.CORSConfig(e)

	routes.GetUserApiRoutes(e, userController)
	routes.GetSwaggerRoutes(e)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}

func init() {
	DB, err := config.PostgresqlConnection()
	if err != nil {
		log.Println("Error when connect postgresql : ", err.Error())
	}

	userRepository := repository.NewUserRepository(DB)
	userController = controller.NewUserController(userRepository)
}
