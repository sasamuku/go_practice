package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db: author` // テーブルのカラム名と異なる場合は明確に指示する
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp4 dbname=gwp4 password=gwp4 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello!", AuthorName: "Hoge"}
	post.Create()
	fmt.Println(post)
}
