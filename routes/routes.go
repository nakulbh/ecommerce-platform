package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nakulbh/ecommerce-platform/controllers"
)

func UserRoutes(e *chi.Mux ){
	e.Post("/users/signup",controllers.Signup())
	e.Post("users/login",controllers.Login())
	e.Post("/admin/addproduct",controllers.ProductViewerAdmin())
	e.Get("/users/productview",controllers.SearchProduct())
	e.Get("users/search",controllers.SearchProductByQuery())

}

func AuthenticationRoutes(e *chi.Mux){
	e.Get("/addtocart", )
}