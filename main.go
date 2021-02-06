package main

import (
	"fmt"
	"net/http"

	"webdev_go/controllers"
	"webdev_go/models"

	"github.com/gorilla/mux"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "db"
)

func main() {

	// Create a DB connection string and then use it to
	// create our model services
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	// =v=v=v= Controllers and Routing =v=v=v=

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()

	// Relatively static pages
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	// =v=v=v= Routed pages =v=v=v=

	// Signing up a user
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}
