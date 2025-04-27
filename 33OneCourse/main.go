package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type Course struct {
	CourseId  string `json:"course_id"`
	CourseName string `json:"course_name"`
	Price   int `json:"price"`
	Author  *Author  `json:"author"` // pointer to Author struct
}

type Author struct {
	FirstName  string  `json:"first_name`
	Website string `json:"website`
}


//Fake Db

var courses []Course

// middleware, helper  - file
func (c *Course) IsEmpty() bool{
	return c.CourseName == "" && c.CourseId == "" && c.Price == 0
}


// controllr = file
 //server  home route 

 func serverHome(w http.ResponseWriter, r http.Request){
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


func main() {
	fmt.Println("Welcome to Api in golang")
}