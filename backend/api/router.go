package api

import (
	"apica-assignment/service"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	cacheService := service.NewCacheService()

	handler := NewHandler(cacheService)

	// Set up CORS middleware to handle preflight requests
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/item", handler.GetItem).Methods("GET")
	router.HandleFunc("/item", handler.SetItem).Methods("POST")
	router.HandleFunc("/item", handler.DeleteItem).Methods("DELETE")

	// Handle OPTIONS for all routes
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Send response to indicate CORS policy is in place
		w.WriteHeader(http.StatusOK)
	})

	return router
}
