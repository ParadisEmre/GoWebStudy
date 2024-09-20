package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ParadisEmre/GoWebStudy/services/product"
	"github.com/ParadisEmre/GoWebStudy/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("Listening on", s.addr)
	log.Println("Process PID", os.Getpid())

	return http.ListenAndServe(s.addr, router)
}
