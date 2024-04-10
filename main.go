package main

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nakulbh/ecommerce-platform/middlewares"
	"github.com/nakulbh/ecommerce-platform/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	r := chi.NewRouter()
	r.Use(middleware.Logger())
	routes.UserRoutes(r)


	r.Use(middlewares.Authentication())
	

	

}