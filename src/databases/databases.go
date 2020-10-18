package databases

import (
	"github.com/AnD00/my-dear-tasks/src/models"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Initialize() {
	// 宣言済みのグローバル変数dbをshort variable declaration(:=)で初期化しようとするとローカル変数dbを初期化することになるので注意する
	db, _ = gorm.Open("sqlite3", "task.db")
	// if err != nil {
	// 	 panic("failed to connect database\n")
	// }

	db.LogMode(true)
	db.AutoMigrate(&models.Task{})
}

func Get() *gorm.DB {
	if db == nil {
		Initialize()
	}
	return db
}

func Close() {
	db.Close()
}
