package main

import (
	"html/template"
	"log"
	"net/http"
)

// go tool dist list => pour avoir la liste de toutes les OS/ARCH
// GOOS=android GOARCH=arm64 go build . => pour lancer le build windows par exemple
// /!\ Les noms des variables dans les structs doivent être en Maj
//
//	pour être exportées dans le html....
type dataIndexHTML struct {
	Titre string
	Nom   string
}

func main() {

	http.HandleFunc("/", HomeHandler)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./public/html"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/html/index.html"))
	data := dataIndexHTML{
		Titre: "Mon titre",
	}
	tmpl.Execute(w, data)
}
