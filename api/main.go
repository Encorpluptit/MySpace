package main

import (
	"MySpace-Api/controllers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	server := new(controllers.Server)
	if godotenv.Load() != nil {
		log.Fatalf("Cannot Load .env file.")
	}
	server.Init()

	defer server.Destroy()
	log.Println("Server runs on http://localhost:" + server.Port)
	log.Fatal(
		http.ListenAndServe(":"+server.Port, server.HandelerCores()(server.Router)))
}
