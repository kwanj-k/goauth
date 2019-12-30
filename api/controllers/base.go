package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/kwanj-k/goauth/api/models"
)

// Server will define our database connection and our routes
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize will have our database connection information, initialise our routes
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// Run will start our server
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	// Use default options
	handler := cors.Default().Handler(server.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
