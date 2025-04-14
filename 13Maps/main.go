package main

import "fmt"

func main(){
	fmt.Println("Welcome to maps of golang");

	lang := make(map[string]string);
	lang["js"] = "javascript";
	lang["py"] = "python";
	lang["rb"] = "ruby";
	lang["go"] = "golang";

	fmt.Println("Languages are : ",lang["js"]);
	delete(lang, "rb");
	fmt.Println("Languages are : ",lang);

	//loops are intertesting in maps

	for key, value := range lang {
		fmt.Printf("For key %v, value is %v\n", key, value);
	}


	languages := make(map[int]string);
	languages[1] = "javascript";
	languages[2] = "python";
	languages[3] = "ruby";
	languages[4] = "golang";
	fmt.Println("Languages are : ",languages);
	fmt.Println("Languages are : ",languages[1]);
	delete(languages, 3);
	fmt.Println("Languages are : ",languages);

	for key, value := range languages{
		fmt.Println("String are : ", value);
		fmt.Printf("For key %v, value is %v\n", key, value);
	}


}