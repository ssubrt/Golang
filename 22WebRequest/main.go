package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var url string = "https://github.com/ssubrt"

func main() {
	fmt.Println("Welcome to 22WebRequest in golang");


	response, err := http.Get(url);

	checkNilError(err); //NOTE : Checking for error

	 fmt.Println("Response code is : %T\n",response.StatusCode); //NOTE : Print the response code

	defer response.Body.Close(); //NOTE : Close the response body after reading

	// databytes , err := ioutil.ReadAll(response.Body); // Method 1
	databytes , err := io.ReadAll(response.Body) // Method 2

	checkNilError(err); //NOTE : Checking for error

	content := string(databytes);
	fmt.Println(content);


}

func checkNilError(err error){
	if err != nil{
		panic(err);
	}
}