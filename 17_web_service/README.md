## PostgreSQL インストール
1. `brew install postgresql`
2. `postgres --version`

## DB セットアップ
1. `createuser -P -d gwp6`
2. `createdb gwp6`
3. `psql -U gwp6 -f setup.sql -d gwp6`

## テスト

新規作成
```
$ curl -i -X POST -H "ContentType: application/json" -d '{"Content":"First Content", "author":"sasamuku"}' http://localhost:8080/post/
HTTP/1.1 200 OK
Date: Sat, 06 Nov 2021 01:02:30 GMT
Content-Length: 0
```

```
$ psql -U gwp6 -d gwp6 -c "select * from posts;"
 id |    content    |  author  
----+---------------+----------
  1 | First Content | sasamuku
(1 row)
```

読み出し
```
$ curl -i -X GET http://localhost:8080/post/1
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 06 Nov 2021 01:06:33 GMT
Content-Length: 67

{
		"id": 1,
		"content": "First Content",
		"author": "sasamuku"
}
```

更新
```
$ curl -i -X PUT -H "Content-Type: application/json" -d '{"Content":"Modified Content", "author":"sasamuku"}' http://localhost:8080/post/1
HTTP/1.1 200 OK
Date: Sat, 06 Nov 2021 01:11:25 GMT
Content-Length: 0
```

```
$ psql -U gwp6 -d gwp6 -c "select * from posts;"
 id |     content      |  author  
----+------------------+----------
  1 | Modified Content | sasamuku
```

削除
```
$ curl -i -X DELETE http://localhost:8080/post/1
HTTP/1.1 200 OK
Date: Sat, 06 Nov 2021 01:13:11 GMT
Content-Length: 0
```

```
$ psql -U gwp6 -d gwp6 -c "select * from posts;"
 id | content | author 
----+---------+--------
(0 rows)
```
