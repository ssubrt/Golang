package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string 
	Price int `json:"itne ka hai bhaiya "`
	Platform string
	password string `json:"-"`
	Tags []string
}


func CheckNilError(err error){
	if err != nil{
		panic(err);
	}
}

func main() {
	fmt.Println("Welcome to JSON in Golang");
	ParseJson();

}

func ParseJson(){
	listofCourses := []course{
		{"ReactJS Bootcamp",199,"coursehub","abc123",[]string {"google","mast"}},
		{"NextJS Bootcamp",399,"coursehub","bcd123",[]string {"golang","hai"}},
		{"NodeJs Bootcamp",299,"coursehub","cde123",nil},
	};

	finalJson , err := json.MarshalIndent(listofCourses,"","\t");
	CheckNilError(err);
	fmt.Printf("Json is : %s\n ",finalJson);
}