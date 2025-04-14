package main

import "fmt"
 
//  checking := "shf"   This is not allowed this iso yṡ allowed within the function Member 
const checking string = "yoyo"  // this allowed at package level and declared as globale variables or values  

func main() {
	var username string = "Subrat"
	fmt.Println((username));
	fmt.Printf("Type of username is %T\n", username);
	var smallVal uint8 = 255
	fmt.Println(smallVal);
	fmt.Printf("Type of username is %T\n", smallVal);
	var isLog bool = false
	fmt.Println(!isLog);
	fmt.Printf("Type of username is %T\n", isLog);
	var smallFloat float64 = 255.45656897
	fmt.Println(smallFloat);
	fmt.Printf("Type of username is %T\n", smallFloat);
	var checkDefaultValue int
	fmt.Println(checkDefaultValue);
	fmt.Printf("Type of username is %T\n", checkDefaultValue);
	var checkDefaultValue2 string
	fmt.Println(checkDefaultValue2 + "1");
	fmt.Printf("Type of username is %T\n", checkDefaultValue2);

	//no var keyword
	NoOfItems := 10
	fmt.Println(NoOfItems);

	fmt.Println(checking);
	fmt.Printf("Type of username is %T\n", checking);
}