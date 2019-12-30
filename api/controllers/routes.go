package controllers

import (
	"github.com/kwanj-k/goauth/api/middlewares"
	"github.com/rs/cors"
)

func (s *Server) initializeRoutes() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	s.Router.StrictSlash(false)

	s.Router.Use(c.Handler)		
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Auth Routes
	s.Router.HandleFunc("/login/", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/users/", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users/",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{id}/",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/users/{id}/",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}/",
		middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
}
