package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	h "github.com/oddball707/recipes/handler"
	s "github.com/oddball707/recipes/service"
)

func main() {
	// initialize service layer
	srv := s.NewRecipeService(db)
	hnd := h.NewRecipeHandler(srv)
	router := hnd.NewRouter()

	mux := http.NewServeMux()

	f := func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	}

	mux.HandleFunc("/", f)

	lambda.Start(httpadapter.New(mux).ProxyWithContext)
}
