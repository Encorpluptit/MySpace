package controllers

import (
	"encoding/json"
	"net/http"
)

// Home : Root endpoint
func Home(w http.ResponseWriter, r *http.Request) {
	if json.NewEncoder(w).Encode("welcome !") != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if json.NewEncoder(w).Encode("world") != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK)
}
