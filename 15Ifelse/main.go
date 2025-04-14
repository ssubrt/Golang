package main

import "fmt"


func main() {
	fmt.Println("Welcome to if else in golang");

	 var result string
	var num int = 10;
	// loginCount := 23;

	if num < 10 {
		result = "less than 10"
	} else if num > 10 {
		result = "greater than 10"
	} else {
		result = "equal to 10"
	}

	fmt.Println(result);


	//NOTE:- In  Web dev we use this type of if else statement
	if num := 3; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("greater than 10")
	}

	var err error  
	if err != nil {
		fmt.Println("Error is : ", err)
	} else {
		fmt.Println("No error")
	}



}