package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"golang-echo-postgresql-rest-api-example/model"
	"log"
)

func PostgresqlConnection() (*gorm.DB, error) {

	Dbdriver := "postgres"

	DB, err := gorm.Open(Dbdriver, PostgresqlUrl)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	DB.AutoMigrate(&model.User{}) //database migration

	return DB, nil
}
