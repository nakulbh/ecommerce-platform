package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nakulbh/ecommerce-platform/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	e := echo.New()
	e.Use(middleware.Logger())
	routes.UserRoutes(e)


	e.Use()
	

	e.Logger.Fatal(e.Start(":" + port))
	

}