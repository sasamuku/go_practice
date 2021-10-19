package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // パッケージを直接使用しない場合アンスコを付ける
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{} // なぜ「:=」としないんだっけ？
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World", Author: "Hoge Hoge"} // Id はDBによって自動付与される
	fmt.Println("Before create:", post)                       // Idの値に注目

	// 投稿作成
	post.Create()
	fmt.Println("Create:", post) // Idの値に注目

	// 投稿取得
	readPost, _ := GetPost(1)
	fmt.Println("Read:", readPost)

	// 投稿更新
	readPost.Content = "Good Morning"
	readPost.Author = "Hoge Hoge Hoge"
	readPost.Update()
	readPost, _ = GetPost(1)
	fmt.Println("Update:", readPost)

	// 投稿削除
	readPost.Delete()
}
