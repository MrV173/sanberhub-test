package database

import (
	"fmt"
	"sanberhub-test/models"
	postgre "sanberhub-test/pkg"
)

func RunMigration() {
	err := postgre.DB.AutoMigrate(
		&models.User{},
		&models.Transaction{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
