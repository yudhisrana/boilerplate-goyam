package main

import (
	"fmt"
	"log"

	"github.com/yudhisrana/boilerplate-goyam/internal/config"
)

func initConfig() (err error) {
	fileName := "config.yaml"
	return config.LoadConfig(fileName)
}

func main() {
	// init config
	if err := initConfig(); err != nil {
		log.Fatalf("\nerror initialize config: %s\n", err.Error())
	}

	// main app
	appName := config.Cfg.App.Name
	fmt.Printf("aplikasi %s berhasil dijalankan\n", appName)
}
