package routes

import (
	"encoding/json"
	"net/http"

	"github.com/stanleyh24/go-gorm-restapi/db"
	"github.com/stanleyh24/go-gorm-restapi/models"

	"github.com/gorilla/mux"
)

func GetTasksHamdler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHamdler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func CreateTaskHamdler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating user: " + err.Error()))
	}
	json.NewEncoder(w).Encode(&task)

}
func DeleteTaskHamdler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}
