package controllers

import (
	"MySpace-Api/database"
	"MySpace-Api/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Port   string
	Router *mux.Router
	DB     database.Database
}

func (s *Server) Init() {
	s.Port = os.Getenv("PORT")
	if s.Port == "" {
		s.Port = "8080"
		log.Printf("Defaulting to port %s", s.Port)
	}

	if err := s.DB.Init(); err != nil {
		log.Fatalf("Database Initialisation Failed: %v", err)
	}

	s.Router = mux.NewRouter().StrictSlash(false)
	s.initialiseRoutes()
}

func (s *Server) Destroy() {
	s.DB.Destroy()
}

func (s *Server) HandelerCores() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))
}

func (s *Server) initialiseRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(Home)).Methods("GET")
	s.Router.HandleFunc("/hello", middlewares.NoCheck(Hello)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

}
