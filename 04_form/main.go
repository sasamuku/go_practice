package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// enctype="application/x-www-form-urlencoded"
	r.ParseForm()
	fmt.Fprintln(w, r.Form)     // Formもクエリ文字列も取得
	fmt.Fprintln(w, r.PostForm) // Formの値のみ取得
	// enctype="multipart/form-data"
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm) // Formの値のみ取得
	// 簡単にFormを取得する方法
	fmt.Fprintln(w, r.FormValue("hello"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
