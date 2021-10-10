package main

import (
	"fmt"
	"net/http"
)

type AppleHandler struct{}

func (h *AppleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Apple!")
}

type OrangeHandler struct{}

func (h *OrangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Orange!")
}

func lemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Lemon!")
}

func grape(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Grape!")
}

func main() {
	apple := AppleHandler{}
	orange := OrangeHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handler未指定の場合、DefaultServeMuxをハンドラとして利用
	}

	// ハンドラによるリクエスト処理
	// ハンドル(appleやorange)をDefaultServeMuxに紐付ける
	http.Handle("/apple", &apple)
	http.Handle("/orange", &orange)

	// ハンドラ関数によるリクエスト処理
	// ハンドラ関数を持ったハンドラに自動変換しDefaultServeMuxに紐付ける
	http.HandleFunc("/lemon", lemon)
	http.HandleFunc("/grape", grape)

	// 既存のインターフェースをハンドラとして使いたいならハンドラ、
	// そうでないなら、ハンドラ関数を使うのがよさそう

	server.ListenAndServe()
}
