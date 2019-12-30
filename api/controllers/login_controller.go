package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kwanj-k/goauth/api/auth"
	"github.com/kwanj-k/goauth/api/models"
	"github.com/kwanj-k/goauth/api/responses"
	"golang.org/x/crypto/bcrypt"
)

// Login func that logins a user and return token
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		var resp = map[string]interface{}{"username": "Wrong email or password"}
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}
	var resp = map[string]interface{}{"status": "Success!"}
	resp["token"] = token
	resp["email"] = user.Email
	responses.JSON(w, http.StatusOK, resp)
}

// SignIn func to verify user details against the db and return token
func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
