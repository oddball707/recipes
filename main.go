package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oddball707/recipes/config"
	d "github.com/oddball707/recipes/dao"
	h "github.com/oddball707/recipes/handler"
	s "github.com/oddball707/recipes/service"
)

func main() {
	// Load configuration
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Error loading config - ", err)
	}

	// Initialize database connection
	db, err := d.NewDatabase(*cfg)
	if err != nil {
		log.Fatal("Error connecting to database - ", err)
	}
	defer db.Close()

	// Initialize DAOs and services
	recipeDAO := d.NewRecipeDAO(db)
	recipeService := s.NewRecipeService(recipeDAO)
	hnd := h.NewRecipeHandler(recipeService)

	router := hnd.NewRouter()
	http.Handle("/", router)

	fmt.Println("Starting server...")
	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	}
}
