package main

import "fmt"

// NOTE:-  No classes in Golang , we have struct in golang
//NOTE:- no inheritamce in Golang


type User struct {
	Name string
	Age int
	Email string
	Status bool
}


func main(){
	fmt.Println("Welcome to structs of golang");

	hitesh := User{Name: "Subrat", Age: 22, Email: "test@gmail.com", Status: true} // Syntax 1
	hitesh2 := User{"Subrat",22,"test@gmail.com",true}; // Syntax 2
	fmt.Println("Hitesh is : ",hitesh);
	fmt.Println("Hitesh2 is : ",hitesh2);
	fmt.Printf("Hitesh2 is : %v\n ",hitesh2);
	fmt.Printf("Hitesh2 is %v and email is %v .",hitesh.Name,hitesh.Email);
}