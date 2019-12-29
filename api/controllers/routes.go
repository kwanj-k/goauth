package controllers

import (
	"github.com/kwanj-k/goauth/api/middlewares"
)

func (s *Server) initializeRoutes() {

	s.Router.StrictSlash(false)

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
