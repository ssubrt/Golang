package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello, World!");
	PerformPostRequest()

}


func CheckNilError(err error){
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error")
	}
}

func PerformPostRequest(){

	const url = "http://localhost:8000/post"


	//Fake Json payload
	requestBody := strings.NewReader(`
	   {
		"course_name" : "Golang",
		"course_price" : 0,
		"platform" : "IPL",
		
		}
	`)

	response , err := http.Post(url, "application/json", requestBody);
	CheckNilError(err);

	defer response.Body.Close()


	content , _ := io.ReadAll(response.Body);
	fmt.Println("Content is :", string(content))

}