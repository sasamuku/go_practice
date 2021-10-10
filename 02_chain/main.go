package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type AppleHandler struct{}

func (h AppleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Apple!")
}

func lemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Lemon!")
}

// ハンドラのチェイン
func log_handler(h http.Handler) http.Handler {
	// ハンドラを生成して返す
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler caller - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

// ハンドラ関数のチェイン
func log_handler_func(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler caller - " + name)
		h(w, r)
	}
}

func main() {
	apple := AppleHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// ハンドラによるリクエスト処理
	http.Handle("/apple", log_handler(apple))

	// ハンドラ関数によるリクエスト処理
	http.HandleFunc("/lemon", log_handler_func(lemon))

	server.ListenAndServe()
}
