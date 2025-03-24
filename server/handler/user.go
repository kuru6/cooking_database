package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleUserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートをパースする
		tpl, err := template.ParseFiles("templates/firstpage.html")
		if err != nil {
			fmt.Println("テンプレート読み込みエラー:", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		//htmlをクライアントに返却
		err = tpl.Execute(c.Response().Writer, nil)
		if err != nil {
			fmt.Println("テンプレート実行エラー:", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		// POST通信のnameを受け取る
		name := c.Request().Header.Get("accountname")
		if name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "トークンが取得できませんでした"})
		}

		// POST通信でpasswordを受け取る
		password := c.Request().Header.Get("password")
		if password == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "トークンが取得できませんでした"})
		}

		fmt.Println(name, password)

		return nil
	}
}

func HandleUserRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートをパースする
		tpl, err := template.ParseFiles("templates/user-register.html")
		if err != nil {
			fmt.Println("テンプレート読み込みエラー:", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		//htmlをクライアントに返却
		err = tpl.Execute(c.Response().Writer, nil)
		if err != nil {
			fmt.Println("テンプレート実行エラー:", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		return nil
	}
}
