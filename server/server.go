package server

import (
	"cooking_database/server/handler"

	"github.com/labstack/echo/v4"
)

// サーバーを起動する関数
func Serve() {
	e := echo.New()

	// 初めのページ兼ログインページ
	e.POST("/login", handler.HandleUserLogin())
	//e.GET("/login", handler.HandleUserLogin())

	// ユーザー登録画面
	e.POST("/user-register", handler.HandleUserRegister())

	// 料理のレシピリストを表示
	e.GET("/recipe-display", handler.HandleRecipeDisplay())

	// 料理のレシピを登録
	e.GET("/recipe-register", handler.HandleRecipeRegister())

	// 料理のレシピを削除
	e.GET("/recipe-delete", handler.HandleRecipeDelete())

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
