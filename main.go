package main

import (
	"fmt"
	"log"
	"test/config"
	router "test/internal/server"

	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := config.MustLoadEnv().Port
	r := router.MainRouter()

	log.Printf("Server is running on http://localhost:%s", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)

}
