package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to 21Files in golang");
	var content string = "This needs to be written in file";

	file, err := os.Create("./mytest.txt")

	checkNilError(err); // check for error

	length , err := io.WriteString(file, content);

	fmt.Println("Length of file is : ", length);
	defer file.Close(); // Close the file after writing to it
	readFile("./mytest.txt") // Read the file after writing to it

}


func readFile(filename string){
	// databyte, err := os.ReadFile(filename); // Method 1
	databyte, err := ioutil.ReadFile(filename);  // Methiod 2
	checkNilError(err);

	fmt.Println("Data in file is : ", databyte);
	fmt.Println("Data in file is : ", string(databyte));
}

func checkNilError(err error){
	if err != nil{
		panic(err);
	}
}