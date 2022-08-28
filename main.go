package main

import (
	"log"
	"net/http"
	"html/template"
)

// go tool dist list => pour avoir la liste de toutes les OS/ARCH
// GOOS=android GOARCH=arm64 go build . => pour lancer le build windows par exemple

 

func main() {
	
	 http.HandleFunc("/", HomeHandler)
	 log.Fatal(http.ListenAndServe(":8000", nil))	  
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w,nil)
}