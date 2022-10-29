package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	model "github.com/hritikkhetan/goRestApis/models"
)

var courses []model.Course

func init() {

	//seeding
	courses = append(courses, model.Course{CourseId: "2", CourseName: "Golang Tutorial", CoursePrice: 499, Author: &model.Author{FullName: "Hitesh Chowdhury", Website: "lco.dev"}})
	courses = append(courses, model.Course{CourseId: "4", CourseName: "Java Tutorial", CoursePrice: 999, Author: &model.Author{FullName: "Telusko", Website: "telusko.com"}})

}

// serve home route
func ServeHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Welcome to the Course Home Page</h1>"))
}

func GetAllCourses(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetCourseById(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get course by id")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(req)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id " + params["id"])
	return
}

func CreateCourse(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if the body is empty
	if req.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course model.Course
	_ = json.NewDecoder(req.Body).Decode(&course)

	// what if the user passes empty json
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send complete data")
	}

	// generating unique id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateCourseById(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Update course by id")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(req)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(req.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id " + params["id"])
	return
}

func DeleteCourseById(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Delete course by id")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(req)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id " + params["id"])
	return
}
