package main

import (
	"database/sql"
	"log"

	"github.com/ParadisEmre/GoWebStudy/cmd/api"
	"github.com/ParadisEmre/GoWebStudy/config"
	"github.com/ParadisEmre/GoWebStudy/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLConnection(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection is successful")
}
