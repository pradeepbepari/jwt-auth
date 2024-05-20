package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/pradeep/golang-micro/cmd/api"
	"github.com/pradeep/golang-micro/config"
	"github.com/pradeep/golang-micro/database"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := database.ConnectionDB(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connection Successful.....")
	server := api.NewApiServer(fmt.Sprintf(":%s", config.Env.Port), db)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
