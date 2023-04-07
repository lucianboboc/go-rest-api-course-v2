package main

import (
	"fmt"

	"github.com/lucianboboc/go-rest-api-course/internal/comment"
	"github.com/lucianboboc/go-rest-api-course/internal/db"
	transportHttp "github.com/lucianboboc/go-rest-api-course/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
