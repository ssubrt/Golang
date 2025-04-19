package main

import (
	"fmt"
	"net/url"
)


var myurl string = "https://localhost:5000?course=react-js"

func checkNilError(err error) {
	if err != nil{
		panic(err);
	}
	
}

func main(){
	
	result , err := url.Parse(myurl);
	checkNilError(err); //NOTE : Checking for error
	fmt.Println("Result is : ", result.Scheme); //NOTE : Print the response code
	fmt.Println("Result is : ", result.Host); //NOTE : Print the response code
	fmt.Println("Result is : ", result.Path); //NOTE : Print the response code
	fmt.Println("Result is : ", result.RawQuery); //NOTE : Print the response code
	fmt.Println("Result is : ", result.Port()); 


	qparams := result.Query();
	fmt.Println("the type of query params is : ", qparams); 
	fmt.Println("the type of query params is : ", qparams["course"]); //NOTE : Print the response code

	for _ ,val := range qparams {
		fmt.Println("Parameter is : ",val);
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=1",

	}

	anotherUrl := partsOfUrl.String();
	fmt.Println("The URL is : ", anotherUrl); //NOTE : Print the response code



}



