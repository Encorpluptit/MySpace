package main

import (
	"MySpace-Api/controllers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func loadEnv() {
	if os.Getenv("API_STATE") != "RELEASE" {
		if godotenv.Load() != nil {
			log.Fatalf("Cannot Load .env file.")
		}
	}
}

//Start Server -> Serve routes -> Defer server destroy
func main() {
	server := new(controllers.Server)
	loadEnv()
	server.Init()

	defer server.Destroy()
	log.Println("Server runs on http://localhost:" + server.Port)
	log.Fatal(
		http.ListenAndServe(":"+server.Port, server.HandelerCores()(server.Router)))
}
