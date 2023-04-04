package main

import (
	"context"
	"fmt"

	"github.com/saifsalah/go_api_v2/internal/comment"
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
	fmt.Println("DB CONNECTED!!")

	if err := db.MigrateDatabase(); err != nil {
		fmt.Println("Fail to migrate the database :( ")
		return err
	}

	commentService := comment.NewService(db)

	commentService.PostComment(context.Background(), comment.Comment{
		Slug:   "saif-test1222",
		Author: "Saif222",
		Body:   "Ilove u golang1222",
	})
	fmt.Println(commentService.GetComment(context.Background(), "6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	return nil

}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
