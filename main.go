package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"test.com/test/router"
)

func main() {
	godotenv.Load()

	router := router.InitRouter()

	http.ListenAndServe(":3000", router)
}
