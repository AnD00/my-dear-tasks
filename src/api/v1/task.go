package v1

import (
	"fmt"
	"net/http"

	"github.com/AnD00/my-dear-tasks/src/enum"
	"github.com/AnD00/my-dear-tasks/src/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ApiTaskHandler struct {
	Db *gorm.DB
}

// 全件取得
// $ curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/ | jsonpp
func (handler *ApiTaskHandler) GetAll(c *gin.Context) {
	var tasks []models.Task      // レコード一覧を格納するため、Task構造体のスライスを変数宣言
	handler.Db.Find(&tasks)      // DBから全てのレコードを取得する
	c.JSON(http.StatusOK, tasks) // JSONで全てのレコードを渡す
}

// 新規作成
// $ curl -X POST -H "Content-Type: application/json" -d '{"text":"test"}' http://localhost:8080/api/v1/ | jsonpp
func (handler *ApiTaskHandler) Create(c *gin.Context) {
	task := models.Task{}
	task.Status = enum.OPEN
	err := c.BindJSON(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	handler.Db.Create(&task)     // レコードを挿入する
	c.JSON(http.StatusOK, &task) // JSONで結果を返す
}

// 取得
// $ curl -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/1 | jsonpp
func (handler *ApiTaskHandler) Get(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // リクエストからidを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	c.JSON(http.StatusOK, task) // JSONで取得したレコードを返す
}

// 更新
// $ curl -X PUT -H "Content-Type: application/json" -d '{"text":"update", "status":2}' http://localhost:8080/api/v1/1 | jsonpp
func (handler *ApiTaskHandler) Update(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // idを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	request := models.Task{}    // Task構造体の変数宣言
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Text = request.Text     // textを上書きする
	task.Status = request.Status // Statusを上書きする
	handler.Db.Save(&task)       // 指定のレコードを更新する
	c.JSON(http.StatusOK, &task) // JSONで結果を返す
}

// 削除
// $ curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/api/v1/1 | jsonpp
func (handler *ApiTaskHandler) Delete(c *gin.Context) {
	task := models.Task{}       // Task構造体の変数宣言
	id := c.Param("id")         // リクエストからidを取得
	handler.Db.First(&task, id) // idに一致するレコードを取得する
	handler.Db.Delete(&task)    // 指定のレコードを削除する
	msg := fmt.Sprintf("Task [%s] has been deleted.", id)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
