package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/guilhermeramos0603/go-api.git/cmd/api"
	"github.com/guilhermeramos0603/go-api.git/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "asd",
		Addr:                 "",
		DBName:               "",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
