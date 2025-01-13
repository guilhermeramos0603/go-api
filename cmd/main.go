package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/guilhermeramos0603/go-api.git/cmd/api"
	"github.com/guilhermeramos0603/go-api.git/config"
	"github.com/guilhermeramos0603/go-api.git/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
