package main

import (
	"log"
	"net/http"

	"github.com/jeauchter/adjutor-auth-api/config"
	"github.com/jeauchter/adjutor-auth-api/db"
	"github.com/jeauchter/adjutor-auth-api/handlers"
	"github.com/jeauchter/adjutor-auth-api/middleware"

	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Print("config: ", config.Database.Host)
	db, err := db.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection initialized")

	r := mux.NewRouter()

	h := handlers.NewHandler(db)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.ValidateJWT)
	api.HandleFunc("/protected", h.ProtectedEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
