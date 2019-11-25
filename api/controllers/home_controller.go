package controllers

import (
	"net/http"

	"github.com/kwanj-k/goauth/api/responses"
)

// Home welcomes us to the API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}