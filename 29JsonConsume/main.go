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
	// ParseJson();
	DecodeJson();

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

func DecodeJson(){
	jsonDatafromWeb := []byte(`
	{
	      "coursename": "ReactJs Bootcamp",
		  "Price": 299,
		  "website": "coursehub",
		  "tags":["web-dev","mast hai"]
	}
	`)

	var lcoCourses course
	checkValid := json.Valid(jsonDatafromWeb);


	if checkValid {
		fmt.Println("Json is Valid");
		json.Unmarshal(jsonDatafromWeb,&lcoCourses);
		fmt.Printf("%#v\n",lcoCourses);
	} else{
		fmt.Println("Json is not valid");
	}


	var myOnlineData map[string] interface{}
	json.Unmarshal(jsonDatafromWeb,&myOnlineData)

	fmt.Printf("%#v\n",myOnlineData);

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v  and the value is %v and tyoe is %T \n",k,v,v);
	}
}