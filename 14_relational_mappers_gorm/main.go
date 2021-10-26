package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    // Gormが自動的に外部キーと想定して関係を作成
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp5 dbname=gwp5 password=gwp5 sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{}) // 構造体にフィールドの変更があるとテーブルを修正する
}

func main() {
	// ポスト作成
	post := Post{Content: "Hello!", Author: "Judy"}
	fmt.Println("Before create:", post)

	// データ追加
	Db.Create(&post)
	fmt.Println("After create:", post)

	// コメント追加
	comment := Comment{Content: "Good post!", Author: "Mike"}
	Db.Model(&post).Association("Comments").Append(comment)

	// ポスト読込
	var readPost Post
	Db.Where("author = ?", "Judy").First(&readPost)
	fmt.Println("Read post:", readPost)

	// コメント読込
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println("Read comment:", comments[0])
}
