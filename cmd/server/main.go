package main

import (
	"fmt"

	"github.com/saifsalah/go_api_v2/internal/db"
)

// Run will be responsible for the init and startup the go app

func Run() error {
	fmt.Println("Starting up application...")

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to the db")
		return err
	}

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	if err := db.MigrateDatabase(); err != nil {
		fmt.Println("Fail to migrate the database :( ")
		return err
	}

	fmt.Println("DB CONNECTED!!")
	return nil

}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
