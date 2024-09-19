package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ParadisEmre/GoWebStudy/service/user"
	"github.com/gorilla/mux"
)

type APISERVER struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APISERVER {
	return &APISERVER{
		address: address,
		db:      db,
	}
}

func (s *APISERVER) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Server is running on", s.address)

	return http.ListenAndServe(s.address, router)
}
