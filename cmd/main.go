package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yudhisrana/boilerplate-goyam/external/database"
	"github.com/yudhisrana/boilerplate-goyam/internal/config"
)

func initConfig() (err error) {
	fileName := "config.yaml"
	return config.LoadConfig(fileName)
}

func initDB() (db *sql.DB, err error) {
	return database.ConnectDB(config.Cfg.DB)
}

func main() {
	// init config
	if err := initConfig(); err != nil {
		log.Fatalf("\nerror initialize config: %s\n", err.Error())
	}

	// init database
	db, err := initDB()
	if err != nil {
		log.Fatalf("error initialize database: %s\n", err.Error())
	}
	defer db.Close()

	// main app
	appName := config.Cfg.App.Name
	fmt.Printf("aplikasi %s berhasil dijalankan\n", appName)
}
