package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path"
)

type db struct {
	name string
	basesql string
}

var (
	pwd, _ = os.Getwd()
	dbName = path.Join(pwd, "db", "blog.db")
)

var (
	InsertPost = db{"insert post", "insert into post(title, summary, content, category_id) values(?, ?, ?, ?)"}
	InsertCate = db{"insert category", "insert into category(name) values(?)"}
)

type DB interface {
	insertInfo() bool
	deleteInfo() bool
	updateInfo() bool
	selectInfo() map[string]string
}

func (d db) InsertOneInfo(data ...interface{}) {
  db, err := sql.Open("sqlite3", dbName)
  checkErr(err)

  stmt, err := db.Prepare(d.basesql)
  checkErr(err)
  fmt.Println(data)
  _, err = stmt.Exec(data...)
  checkErr(err)
}

func (d db) InsertMoreInfo(data ...[]interface{}) {
	db, err := sql.Open("sqlite3", dbName)
	checkErr(err)
	stmt, err := db.Prepare(d.basesql)
	checkErr(err)
	for _, d := range data {
		_, err = stmt.Exec(d...)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func InitDB() {
	insertCateDB := InsertCate
	insertCateDB.InsertOneInfo("Python")
	insertCateDB.InsertOneInfo("Linux")
	insertCateDB.InsertOneInfo("Golang")

	insertPostDB := InsertPost
	insertPostDB.InsertOneInfo("test_title", "test_summary", "test_content", 1)
}
