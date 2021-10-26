package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // パッケージを直接使用しない場合アンスコを付ける
)

type Post struct {
	Id      int
	Content string
	Author  string
	// 構造体Commentのスライス
	// スライスは配列へのポインタであるため実際にはCommentのポイントである
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp3 dbname=gwp3 password=gwp3 sslmode=disable")
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

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	// Scan()でクエリにより返ってきた値を変数にコピーする
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	if err != nil {
		fmt.Println("Error while creating comment:\n", err)
	}
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		fmt.Println("Error while getting post:\n", err)
		return
	}

	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		fmt.Println("Error while getting comment:\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			fmt.Println("Error while getting each comment:\n", err)
			return
		}
		post.Comments = append(post.Comments, comment)
	}
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
	post := Post{Content: "This is post", Author: "Hoge"} // Id はDBによって自動付与される
	fmt.Println("Before create:", post)                   // Idの値に注目

	// 投稿作成
	post.Create()
	fmt.Println("Create Post:", post)

	// コメント作成
	comment := Comment{Content: "LGMT!", Author: "Taro", Post: &post}
	fmt.Println("Create Comment:", comment)
	comment.Create()

	// 投稿取得
	readPost, _ := GetPost(post.Id)
	fmt.Println("Read Post:", readPost)
	fmt.Println("Read Comments:", readPost.Comments)
	fmt.Println("Read Comment:", readPost.Comments[0].Post)
}
