package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"movie-gopher/handler"
	"net/http"
)

// Think of handler as controller in MVC frameworks. Every controller has controller methods for list, retrieve, update,
// create and delete. Note that I am using Django convention here. List is equivalent to index in Rails; retrieve is
// equivalent to show in Rails.
func LoadRoutes(db *gorm.DB) http.Handler {
	// Defining middleware
	logMiddleware := NewServerLoggingMiddleware()

	muxRouter := mux.NewRouter().StrictSlash(true)

	// Name-spacing the API
	api := muxRouter.PathPrefix("/api").Subrouter()
	api.Handle("/users", handler.NewUserListHandler(db)).Methods("GET")
	api.Handle("/users/{id:[0-9]+}", handler.NewUserRetrieveHandler(db)).Methods("GET")
	api.Handle("/users", handler.NewUserCreateHandler(db)).Methods("POST")
	api.Handle("/authenticate", handler.NewSessionCreateHandler(db)).Methods("POST")

	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	return handlers.CORS()(logMiddleware(muxRouter))
}
