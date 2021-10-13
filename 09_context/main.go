package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func context(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println("Can not load template file")
	}
	content := `This is: <i>"Apple"</i>`
	t.Execute(w, content)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/context", context)
	server.ListenAndServe()
}
