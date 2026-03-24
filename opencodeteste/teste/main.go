package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
	Items   []string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)

	fmt.Println("Servidor rodando em: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := PageData{
		Title:   "Minha Página Go",
		Message: "Olá! Esta página foi criada com Go puro, sem HTML manual!",
		Items:   []string{"Item 1 - Go é poderoso", "Item 2 - Sem React necessário", "Item 3 - Hot reload funcionando"},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
