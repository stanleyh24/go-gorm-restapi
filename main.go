package main

import (
	"net/http"

	"github.com/stanleyh24/go-gorm-restapi/db"
	"github.com/stanleyh24/go-gorm-restapi/models"
	"github.com/stanleyh24/go-gorm-restapi/routes"

	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	//users
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//task
	r.HandleFunc("/task", routes.GetTasksHamdler).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHamdler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.GetTaskHamdler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHamdler).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
