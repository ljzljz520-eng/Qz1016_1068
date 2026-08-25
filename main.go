package main

import (
	"fmt"
	"log"
	"net/http"
	"storeinspection/api"
	"storeinspection/config"
	"storeinspection/service"
	"storeinspection/storage"
)

func main() {
	cfg := config.Default()
	db, err := storage.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	svc := service.New(db)
	mux := api.New(svc)
	fmt.Printf("inspection service listening on %s\n", cfg.ListenAddress)
	log.Fatal(http.ListenAndServe(cfg.ListenAddress, mux))
}
