// リスト7.10
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	fmt.Println(jsonData)         // バイト配列
	fmt.Println(string(jsonData)) // 文字列に変換

	var post Post
	json.Unmarshal(jsonData, &post) // jsonデータを構造体にマッピング
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
	fmt.Println(post.Comments[1].Id)
	fmt.Println(post.Comments[1].Content)
	fmt.Println(post.Comments[1].Author)

}
