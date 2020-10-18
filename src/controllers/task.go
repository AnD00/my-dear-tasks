package controllers

import (
	"net/http"
	"strconv"

	"github.com/AnD00/my-dear-tasks/src/enum"
	"github.com/AnD00/my-dear-tasks/src/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TaskHandler struct {
	Db *gorm.DB
}

// 一覧表示
func (handler *TaskHandler) GetAll(c *gin.Context) {
	var tasks []models.Task                                    // レコード一覧を格納するため、Task構造体のスライスを変数宣言
	handler.Db.Find(&tasks)                                    // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "index.html", gin.H{"tasks": tasks}) // index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *TaskHandler) Create(c *gin.Context) {
	text, _ := c.GetPostForm("text")                               // index.htmlからtextを取得
	handler.Db.Create(&models.Task{Text: text, Status: enum.OPEN}) // レコードを挿入する
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 編集画面
func (handler *TaskHandler) Edit(c *gin.Context) {
	task := models.Task{}                                   // Task構造体の変数宣言
	id := c.Param("id")                                     // index.htmlからidを取得
	handler.Db.First(&task, id)                             // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "edit.html", gin.H{"task": task}) // edit.htmlに編集対象のレコードを渡す
}

// 更新
func (handler *TaskHandler) Update(c *gin.Context) {
	task := models.Task{}                // Task構造体の変数宣言
	id := c.Param("id")                  // edit.htmlからidを取得
	text, _ := c.GetPostForm("text")     // edit.htmlからtextを取得
	status, _ := c.GetPostForm("status") // edit.htmlからstatus(文字列 1~3)を取得
	handler.Db.First(&task, id)          // idに一致するレコードを取得する
	task.Text = text                     // textを上書きする
	num, _ := strconv.Atoi(status)       // 数値に変換
	task.Status = enum.State(num)        // statusを上書きする
	handler.Db.Save(&task)               // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 削除
func (handler *TaskHandler) Delete(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // index.htmlからidを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	handler.Db.Delete(&task)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/")
}
