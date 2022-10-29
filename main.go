package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/hritikkhetan/goRestApis/controllers"
)

// Main function
func main() {

	router := mux.NewRouter()

	// handle routing
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", controller.GetCourseById).Methods("GET")
	router.HandleFunc("/course", controller.CreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}", controller.UpdateCourseById).Methods("PUT")
	router.HandleFunc("/course/{id}", controller.DeleteCourseById).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", router))

}
