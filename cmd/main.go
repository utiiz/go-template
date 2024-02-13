package main

import (
	database "github.com/utiiz/go/templ/internal/db"
	"github.com/utiiz/go/templ/internal/routes"
)

func main() {
	db := database.DB{}

	db.Init()

	router := routes.NewRouter(db)
	router.Logger.Fatal(router.Start(":8080"))
}
