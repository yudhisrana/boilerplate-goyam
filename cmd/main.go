package main

import (
	"fmt"
	"log"

	"github.com/yudhisrana/boilerplate-goyam/internal/config"
)

func main() {
	// setup config
	fileName := "config.yaml"
	if err := config.LoadConfig(fileName); err != nil {
		log.Fatalf("\nerror initialize config: %s\n", err.Error())
	}

	// main app
	appName := config.Cfg.App.Name
	fmt.Printf("aplikasi %s berhasil dijalankan\n", appName)
}
