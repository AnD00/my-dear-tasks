package routes

import (
	v1 "github.com/AnD00/my-dear-tasks/src/api/v1"
	"github.com/AnD00/my-dear-tasks/src/controllers"
	db "github.com/AnD00/my-dear-tasks/src/databases"

	"github.com/gin-gonic/gin"
)

func Router() {
	// gin内で定義されているEngine構造体インスタンスを取得
	// Router、HTML Rendererの機能を内包している
	router := gin.Default()

	// globパターンに一致するHTMLファイルをロードしHTML Rendererに関連付ける
	// 今回のケースではviewsディレクトリ配下のhtmlファイルを関連付けている
	router.LoadHTMLGlob("views/*.html")

	// TaskHandler構造体に紐付けたCRUDメソッドを呼び出す
	handler := controllers.TaskHandler{
		db.Get(),
	}

	// 各パスにGET/POSTメソッドでリクエストされた時のレスポンスを登録する
	// 第一引数にパス、第二引数にHandlerを登録する
	router.GET("/", handler.GetAll)            // 一覧表示
	router.POST("/", handler.Create)           // 新規作成
	router.GET("/edit/:id", handler.Edit)      // 編集画面
	router.POST("/update/:id", handler.Update) // 更新
	router.POST("/delete/:id", handler.Delete) // 削除

	// 共通の /api/v1 をパスに持つルートをグループ化する
	apiV1 := router.Group("/api/v1")
	{
		api := v1.ApiTaskHandler{
			Db: db.Get(),
		}
		// 編集画面の /:id というパスは :id をワイルドカードとして http://localhost:8080/1 だけではなく
		// http://localhost:8080/api/vi/ も同一パスとして判定してしまう
		// そのため /edit/:id に変更した
		// panic: 'api' in new path '/api/v1/' conflicts with existing wildcard ':id' in existing prefix '/:id'
		apiV1.GET("/", api.GetAll)       // 全件取得
		apiV1.POST("/", api.Create)      // 新規作成
		apiV1.GET("/:id", api.Get)       // 取得
		apiV1.PUT("/:id", api.Update)    // 更新
		apiV1.DELETE("/:id", api.Delete) // 削除
	}

	// Routerをhttp.Serverに接続し、HTTPリクエストのリスニングとサービスを開始する
	router.Run()
}
