package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/nakulbh/ecommerce-platform/controllers"
)

func UserRoutes(e *echo.Echo ){
	e.POST("/users/signup",controllers.Signup())
	e.POST("users/login",controllers.Login())
	e.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	e.GET("/users/productview",controllers.SearchProduct())
	e.GET("users/search",controllers.SearchProductByQuery())

}