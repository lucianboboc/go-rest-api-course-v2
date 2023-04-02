package main

import (
	"context"
	"fmt"

	"github.com/lucianboboc/go-rest-api-course/internal/comment"
	"github.com/lucianboboc/go-rest-api-course/internal/db"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("staring up our application")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err = db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	fmt.Println("Successfully connected and pinged database")

	cmtService := comment.NewService(db)

	cmtService.Store.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "d5689f06-4d14-4944-9810-769e48832949",
			Slug:   "manual-test",
			Author: "Lucian",
			Body:   "Hello World",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"c5689f06-4d14-4944-9710-769e49932944",
	))

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
