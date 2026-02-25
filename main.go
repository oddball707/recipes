package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	d "github.com/oddball707/recipes/dao"
	h "github.com/oddball707/recipes/handler"
	s "github.com/oddball707/recipes/service"
)

func main() {
	// Load configuration
	cfg, err := LoadConfig()
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

	serve(recipeService)
}

func serve(srv s.RecipeClient) {
	hnd := h.NewRecipeHandler(srv)
	router := hnd.NewRouter()

	http.Handle("/", router)
	port := getEnv("PORT", "8080")

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// cannot panic, because this probably is an intentional close
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadConfig() (*d.DBConfig, error) {
	cfg := &d.DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "recipes"),
	}

	return cfg, nil
}
