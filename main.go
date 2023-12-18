package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", add)
	http.HandleFunc("/view", view)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)

}

func add(w http.ResponseWriter, r *http.Request){

	ptmp,err := template.ParseFiles("template/base.html")
	if err!=nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	// fmt.Fprintf(w, `welcome`)
}
func view(w http.ResponseWriter, r *http.Request){

	ptmp,err := template.ParseFiles("template/base.html")
	if err!=nil {
		fmt.Println(err.Error())
	}
	ptmp,err = ptmp.ParseFiles("wpage/view.html")
	if err!=nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	// fmt.Fprintf(w, `welcome`)
}





	
	// contact := make(map[string]string)

	// //adding contact
	// contact["name"]= "ami"
	// contact["number"] = "01797185909"
	// fmt.Println("Name:",contact["name"],"\n","Number:",contact["number"])

	// //delete contact
	// delete(contact,"name")
	// delete(contact,"number")

	
	// fmt.Println("Name:",contact["name"],"\n","Number:",contact["number"])


	// var name string
	// var phone int

	// //create contact list
	// fmt.Printf("Name: ")  
	// fmt.Scanf("%s \n",&name)
	// fmt.Printf("Phone Number: ")
	// fmt.Scanf("%d",&phone)


	// //read contact
	// fmt.Println(name, phone)

	


