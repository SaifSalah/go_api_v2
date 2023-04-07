package main

import (
	"fmt"

	"github.com/saifsalah/go_api_v2/internal/comment"
	"github.com/saifsalah/go_api_v2/internal/db"
	transportHttp "github.com/saifsalah/go_api_v2/internal/transport/http"
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
	fmt.Println("DB CONNECTED!!")

	if err := db.MigrateDatabase(); err != nil {
		fmt.Println("Fail to migrate the database :( ")
		return err
	}

	commentService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(commentService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil

}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
