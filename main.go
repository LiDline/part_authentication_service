package main

import (
	"fmt"
	"log"
	"test/config"
	"test/internal/db"
	router "test/internal/server"

	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.MustLoadEnv()

	db.Init(cfg.DbUrl)
	defer db.Close()

	r := router.MainRouter()

	log.Printf("Server is running on http://localhost:%s", cfg.Port)

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)

}
