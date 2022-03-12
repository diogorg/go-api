package routes

import (
	handlers "api/routes/handlers"
	userRoutes "api/routes/usersRoutes"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	userRoutes.Route()
	return handlers.R
}
