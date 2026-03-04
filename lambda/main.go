package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/oddball707/recipes/config"
	"github.com/oddball707/recipes/dao"
	d "github.com/oddball707/recipes/dao"
	h "github.com/oddball707/recipes/handler"
	s "github.com/oddball707/recipes/service"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Error loading config - ", err)
	}

	db, err := d.NewDatabase(*cfg)
	if err != nil {
		log.Fatal("Error connecting to database - ", err)
	}
	defer db.Close()

	dao := dao.NewRecipeDAO(db)
	srv := s.NewRecipeService(dao)
	hnd := h.NewRecipeHandler(srv)
	router := hnd.NewRouter()

	mux := http.NewServeMux()

	f := func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	}

	mux.HandleFunc("/", f)

	lambda.Start(httpadapter.New(mux).ProxyWithContext)
}
