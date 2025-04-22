package main

import "fmt"


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



func main() {
	fmt.Println("Welcome to Api in golang")
}