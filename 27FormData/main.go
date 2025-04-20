package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Hello, World!");
	PerformFormRequest();

}


func CheckNilError(err error){
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error")
	}
}

func PerformFormRequest(){
	const myurl = "http://localhost:8000/post"

	data := url.Values{}
	data.Add("course_name", "Golang");
	data.Add("last_name","Go");
	data.Add("email", "test@go.dev");

	response, err := http.PostForm(myurl, data);
	CheckNilError(err);
	defer response.Body.Close();
	content, err := io.ReadAll(response.Body)
	CheckNilError(err)
	fmt.Println("Content is :", string(content))
}

