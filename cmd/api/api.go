package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matimortari/go-ecom-backend/services/cart"
	"github.com/matimortari/go-ecom-backend/services/order"
	"github.com/matimortari/go-ecom-backend/services/product"
	"github.com/matimortari/go-ecom-backend/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// Create a new APIServer struct
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run the API server
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Server is running on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
