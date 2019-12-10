package db

import (
	"database/sql"
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
	sqldriver = "sqlite3"
	pwd, _ = os.Getwd()
	dbName = path.Join(pwd, "db", "blog.db")
)

var (
	InsertPost = db{"insert post", "insert into post(title, summary, content, category_id) values(?, ?, ?, ?)"}
	InsertCate = db{"insert category", "insert into category(name) values(?)"}
	CreatePostTable = db{"create post table", "CREATE TABLE `post` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT, `create_time` TIMESTAMP default (datetime('now', 'localtime'))," +
        "`update_time` TIMESTAMP default (datetime('now', 'localtime')), `title` VARCHAR(64) NOT NULL, " +
		"`summary` VARCHAR(100) NULL, `content` TEXT NOT NULL, `category_id` INTEGER NOT NULL," +
		"FOREIGN KEY (category_id) REFERENCES category(id));"}
	CreateCateTable = db{"create category table", "CREATE TABLE `category` (`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`create_time` TIMESTAMP default (datetime('now', 'localtime')), " +
		"`update_time` TIMESTAMP default (datetime('now', 'localtime')), `name` VARCHAR(50) NOT NULL);"}
)

type DB interface {
	insertInfo() bool
	deleteInfo() bool
	updateInfo() bool
	selectInfo() map[string]string
}

func CreateDB() error {
	if _, err := os.Stat(dbName); err != nil {
		err = createDBFile()
		if err != nil {
			log.Fatal(err)
			return err
		}
		err = createTable()
		checkErr(err)
	} else {
		err = createTable()
		checkErr(err)
	}
	return nil
}

func createDBFile() error {
	f, err := os.Create(dbName)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = os.Chmod(dbName, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func createTable() error {
	db, err := sql.Open(sqldriver, dbName)
	defer db.Close()
	checkErr(err)

	res, err := db.Exec(CreateCateTable.basesql)
	log.Println("创建分类表：", res)
	checkErr(err)

	res, err = db.Exec(CreatePostTable.basesql)
	log.Println("创建文章表", res)
	checkErr(err)
	return nil
}

func (d db) InsertOneInfo(data ...interface{}) {
  db, err := sql.Open(sqldriver, dbName)
  defer db.Close()
  checkErr(err)

  stmt, err := db.Prepare(d.basesql)
  checkErr(err)
  log.Println("正在初始化数据：", data)
  _, err = stmt.Exec(data...)
  checkErr(err)
}

func (d db) InsertMoreInfo(data ...[]interface{}) {
	db, err := sql.Open("sqlite3", dbName)
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare(d.basesql)
	defer stmt.Close()
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

func RecoverEnv() error {
	if _, err := os.Stat(dbName); !os.IsNotExist(err) {
		err = os.Remove(dbName)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	log.Println("环境处理完成")
	return nil
}
