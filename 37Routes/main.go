package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)


type Course struct {
	CourseId  string `json:"course_id"`
	CourseName string `json:"course_name"`
	Price   int `json:"price"`
	Author  *Author  `json:"author"` // pointer to Author struct
}

type Author struct {
	FirstName  string  `json:"first_name"`
	Website string `json:"website"`
}


//Fake Db

var courses []Course

// middleware, helper  - file
func (c *Course) IsEmpty() bool{
	// return c.CourseName == "" && c.CourseId == "" && c.Price == 0
	return c.CourseName == ""
}


// controllr = file
 //server  home route 

 func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to golang </h1>"))
 }

 func getAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
 }

 func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r);

	fmt.Println(params)

	for _,course := range courses{
		if course.CourseId == params["course_id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(`No course found with given id  : %T` + params["course_id"])
 }

 func createCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data ")
		return
	}


	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No Data inside json")
		return
	}


	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
 }

 func updateCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json");

	// grabbing the id from the request

	params := mux.Vars(r);

	// loop through the courses, find matching id and remove it and then add the new course id
	for index, course := range courses{
		if course.CourseId == params["course_id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["course_id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
		// TODO: send a response when the course is not found
		json.NewEncoder(w).Encode("No course found with given id  : %T" + params["course_id"])
	}
 }


 func deleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json");

	// grabbing the id from the request
	params := mux.Vars(r);


	for index,course := range courses{
		if course.CourseId == params["course_id"]{
			courses = append(courses[:index], courses[index+1:]...);
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}

	}
	json.NewEncoder(w).Encode("No course found with given id  : %T" + params["course_id"])
 }





func main() {
	fmt.Println("Welcome to Api in golang")

	r := mux.NewRouter();

	courses = append(courses, Course{
		CourseId: "1",
		CourseName : "React JS",
		Price: 299,
		Author : &Author{
			FirstName: "John",
			Website: "john.com"},
	})
	courses = append(courses, Course{
		CourseId: "2",
		CourseName : "Next JS",
		Price: 399,
		Author : &Author{
			FirstName: "John",
			Website: "go.com"},
	})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET");
	r.HandleFunc("/courses", getAllCourses).Methods("GET");
	r.HandleFunc("/course/{course_id}", getOneCourse).Methods("GET");
	r.HandleFunc("/course", createCourse).Methods("POST");
	r.HandleFunc("/course/{course_id}", updateCourse).Methods("PUT");
	r.HandleFunc("/course/{course_id}", deleteCourse).Methods("DELETE");



	log.Fatal(http.ListenAndServe(":4000",r));



	// listening to the port
	log.Fatal(http.ListenAndServe(":4000",r));
}