package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder
var post_id int

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	// テストケースで用いるグローバル関数を初期化
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
	post_id = 2
}

func TestHandleGet(t *testing.T) {
	path := "/post/" + strconv.Itoa(post_id)
	request, _ := http.NewRequest("GET", path, nil) // method, path, request body
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != post_id {
		t.Errorf("Cannot retriece JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	body := strings.NewReader(`{"content":"Updated post", "author":"sasamuku2"}`)
	path := "/post/" + strconv.Itoa(post_id)
	request, _ := http.NewRequest("PUT", path, body) // method, path, request body
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
