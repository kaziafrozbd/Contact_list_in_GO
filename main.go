package main

import (
	"fmt"
	"html/template"
	"net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err error
func init(){

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err = sql.Open("mysql", "root:contact123@tcp(127.0.0.1:3306)/contact_list_db")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    // defer db.Close()

	
    


	fmt.Println("database connected")
}



func main() {

	http.HandleFunc("/", add)
	http.HandleFunc("/req", req)
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

func req(w http.ResponseWriter, r *http.Request){

	name := r.FormValue("name")
	phone := r.FormValue("phone")



	qs := "INSERT INTO `contact` (`id`, `name`, `phone`, `status`) VALUES (NULL, '%s', '%s', '1');"
	sql := fmt.Sprintf(qs,name,phone)
	// perform a db.Query insert
    insert, err := db.Query(sql)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
	
	fmt.Println(name,phone)

	fmt.Fprintf(w, `recieved`)
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

	


