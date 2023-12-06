package main

import "fmt"

func main() {
	
	// contact := make(map[string]string)

	// //adding contact
	// contact["name"]= "ami"
	// contact["number"] = "01797185909"
	// fmt.Println("Name:",contact["name"],"\n","Number:",contact["number"])

	// //delete contact
	// delete(contact,"name")
	// delete(contact,"number")

	
	// fmt.Println("Name:",contact["name"],"\n","Number:",contact["number"])


	var name string
	var phone int

	//create contact list
	fmt.Printf("Name: ")  
	fmt.Scanf("%s \n",&name)
	fmt.Printf("Phone Number: ")
	fmt.Scanf("%d",&phone)


	//read contact
	fmt.Println(name, phone)

	


}