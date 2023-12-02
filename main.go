package main

import "fmt"

func main() {
	
	contact := make(map[string]string)
	contact["name"]= "ami"
	contact["number"] = "01797185909"
	fmt.Println("Name:",contact["name"],"\n","Number:",contact["number"])

	delete(contact,"name")
	delete(contact,"number")

	
	fmt.Println("Name:",contact["name"],"\n","Number:",contact["number"])

}