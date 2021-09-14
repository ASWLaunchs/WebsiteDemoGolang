package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./www/index.html")
	if err != nil {
		log.Fatal("template.ParseFiles : ", err)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.Handle("/img/", http.FileServer(http.Dir("www")))
	http.Handle("/js/", http.FileServer(http.Dir("www")))
	http.Handle("/images/", http.FileServer(http.Dir("www")))
	http.Handle("/fonts/", http.FileServer(http.Dir("www")))
	http.Handle("/elegant_font/", http.FileServer(http.Dir("www")))
	http.Handle("/css/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/", handler)
	fmt.Println("server is running in 80 port.")
	log.Fatal(http.ListenAndServe(":80", nil))
}
