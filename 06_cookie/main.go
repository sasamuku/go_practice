package main

import (
	"fmt"
	"net/http"
)

func set_cookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "cookie1",
		Value:    "value of cookie1",
		HttpOnly: true,
	}
	cookie2 := http.Cookie{
		Name:     "cookie2",
		Value:    "value of cookie2",
		HttpOnly: true,
	}
	cookie3 := http.Cookie{
		Name:     "cookie3",
		Value:    "value of cookie3",
		HttpOnly: true,
	}
	cookie4 := http.Cookie{
		Name:     "cookie4",
		Value:    "value of cookie4",
		HttpOnly: true,
	}
	// 値渡し
	w.Header().Set("Set-Cookie", cookie1.String())
	w.Header().Add("Set-Cookie", cookie2.String())

	// 参照渡し
	http.SetCookie(w, &cookie3) // SetCookieの場合は値が追加されるがHeader().Setは値を上書き
	http.SetCookie(w, &cookie4)
}

func get_cookie(w http.ResponseWriter, r *http.Request) {
	cookie1, err := r.Cookie("cookie1")
	if err != nil {
		fmt.Fprintln(w, "Can not get the specified cookie")
	}
	all_cookies := r.Cookies()
	fmt.Fprintln(w, cookie1)
	fmt.Fprintln(w, all_cookies)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/set_cookie", set_cookie)
	http.HandleFunc("/get_cookie", get_cookie)
	server.ListenAndServe()
}
