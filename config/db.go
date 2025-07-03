package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connect := "host=localhost port=5432 user=postgres password=root dbname=postgres sslmode=disable"
	DB, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal melakukan ping ke database:", err)
	}
	log.Println("Berhasil terkoneksi ke database")
}
