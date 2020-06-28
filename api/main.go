package main

import (
	"MySpace-Api/controllers"
	"log"
	"net/http"
)

func main() {
	server := new(controllers.Server)
	server.Init()

	defer server.Destroy()
	log.Println("Server runs on http://localhost:" + server.Port)
	log.Fatal(
		http.ListenAndServe(":"+server.Port, server.HandelerCores()(server.Router)))
}
