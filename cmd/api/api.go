package api

import "database/sql"

type APISERVER struct {
	address string
	db      *sql.DB
}

type NewAPIServer(address string, db *sql.DB) *APISERVER {

}