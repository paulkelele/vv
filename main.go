package main

import (
 	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// go tool dist list => pour avoir la liste de toutes les OS/ARCH
// GOOS=android GOARCH=arm64 go build . => pour lancer le build windows par exemple
// /!\ Les noms des variables dans les structs doivent être en Maj
//
//	pour être exportées dans le html....
type dataIndexHTML struct {
	Titre string
 }
 
 
 
 func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/html/index.html"))
	w.Header().Set("Content-Type", "text/html")
	data := dataIndexHTML{
		Titre: "Mon titre",
	}
	tmpl.Execute(w, data)
}

 

func main() {
	
  	r := mux.NewRouter()
	// page racine
	r.HandleFunc("/", HomeHandler).Methods(http.MethodGet).Methods(http.MethodGet)


	  r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	  r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
	  r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))

 

srv := &http.Server{
	Addr: "0.0.0.0:8000",
	WriteTimeout: time.Second * 15,
	ReadTimeout:  time.Second * 15,
	IdleTimeout:  time.Second * 60,
	Handler: r,
}
log.Fatal(srv.ListenAndServe())
	 
	
}


 