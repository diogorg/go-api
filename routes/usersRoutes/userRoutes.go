package userRoutes

import (
	userController "api/controllers/users"
	routes "api/routes/handlers"
)

var routeBase string = "/users"

func Route() {
	routes.Get(routeBase+"", userController.Index)
	routes.Get(routeBase+"/{id}", userController.Show)
	routes.Post(routeBase+"", userController.Store)
	routes.Delete(routeBase+"/{id}", userController.Delete)
	routes.Patch(routeBase+"/{id}", userController.Update)
}
