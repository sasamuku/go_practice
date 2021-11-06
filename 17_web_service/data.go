package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp6 dbname=gwp6 password=gwp6 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	// QueryRow()はRow型(1行のSQL結果)を返すので、Scan()で変数にそれぞれ割り当てる
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement) // Prepare()は後のクエリのための準備済ステートメントを作成する
	if err != nil {
		return
	}
	defer stmt.Close() // 準備済ステートメントはサーバリソースを確保するため使用後はクローズする
	// QueryRow()は引数を受け取って準備済ステートメントを実行する
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	// 結果を返却せずにクエリを実行する
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
