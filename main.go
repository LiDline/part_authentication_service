package main

import (
	"fmt"
	"test/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	сfg := config.MustLoadEnv()

	fmt.Println(сfg)

}
