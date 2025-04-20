package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello, World!");
	performGetResponse();

}


func CheckNilError(err error){
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error")
	}
}

func performGetResponse() {
	const url = "http://localhost:8000/get"
	
	response, err := http.Get(url);
	CheckNilError(err)

	defer response.Body.Close();

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response length:", response.ContentLength);

	content, _ := io.ReadAll(response.Body);

	var responseString strings.Builder
	byteCount, err := responseString.Write(content);

	fmt.Println("Byte Count:", byteCount);
	fmt.Println(responseString.String());
	fmt.Println("Response Body:", string(content))
}