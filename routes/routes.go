package routes

import (
	"inventory-management-system/controllers"
	"inventory-management-system/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/items", controllers.CreateItemHandler).Methods("POST")
	api.HandleFunc("/items", controllers.ListItemsHandler).Methods("GET")
	api.HandleFunc("/items/{id}", controllers.GetItemDetailHandler).Methods("GET")
	api.HandleFunc("/items/{id}/stock", controllers.UpdateStockHandler).Methods("PATCH")
	api.HandleFunc("/logs", controllers.ListLogsHandler).Methods("GET")
}
