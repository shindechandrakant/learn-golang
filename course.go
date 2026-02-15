package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake Db
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func coursesMain() {
	r := mux.NewRouter()
	// Seeding
	courses = []Course{
		{
			CourseId:    "C101",
			CourseName:  "Go Fundamentals",
			CoursePrice: "1999",
			Author: &Author{
				FullName: "Alice Johnson",
				Website:  "https://alice.dev",
			},
		},
		{
			CourseId:    "C102",
			CourseName:  "Advanced Go Concurrency",
			CoursePrice: "2999",
			Author: &Author{
				FullName: "Bob Smith",
				Website:  "https://bobsmith.io",
			},
		},
		{
			CourseId:    "C103",
			CourseName:  "REST APIs with Go",
			CoursePrice: "2499",
			Author: &Author{
				FullName: "Carol Williams",
				Website:  "https://carol.codes",
			},
		},
		{
			CourseId:    "C104",
			CourseName:  "Microservices in Go",
			CoursePrice: "3499",
			Author: &Author{
				FullName: "David Brown",
				Website:  "https://davidbrown.tech",
			},
		},
		{
			CourseId:    "C105",
			CourseName:  "Go for Backend Engineers",
			CoursePrice: "1799",
			Author: &Author{
				FullName: "Eva Miller",
				Website:  "https://evamiller.dev",
			},
		},
	}

	r.HandleFunc("/", serveHome).Methods("GET", "POST")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	fmt.Println("Server is running")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running")
}

// controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang series..</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop through courses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body can't be empty")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("course is empty")
		return
	}

	course.CourseId = generateUniqueIdId()
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func generateUniqueIdId() string {
	rand.Seed(time.Now().UnixNano())
	courseId := strconv.Itoa(rand.Intn(100))
	return courseId
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for indx, val := range courses {
		if val.CourseId == params["id"] {
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses[:indx], courses[indx+1:]...)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("course updated")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given ID to update")
}
