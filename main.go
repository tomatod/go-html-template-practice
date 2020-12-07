package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
)

func main() {
	showTemplate()
	server()
}

func showTemplate() {
	t, _ := template.ParseFiles("main.html", "header.html", "body.html")

	b := bytes.NewBuffer([]byte{})
	t.Execute(b, struct {
		Message string
	}{
		"hello!",
	})
	b.WriteTo(os.Stdout)
}

func server() {
	http.HandleFunc("/a", templateHandle)
	http.HandleFunc("/b", templateAllInAFileHandle)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.ListenAndServe(":8080", nil)
}

func templateHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html", "header.html", "body.html")
	t.Execute(w, struct {
		Message string
	}{
		"hello!",
	})
}

func templateAllInAFileHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("allInAFile.html")
	t.Execute(w, struct {
		Message string
	}{
		"All in a file!",
	})
}
