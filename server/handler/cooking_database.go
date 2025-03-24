package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleRecipeDisplay() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートをパースする
		tpl, err := template.ParseFiles("templates/recipe-display.html")
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

func HandleRecipeRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートをパースする
		tpl, err := template.ParseFiles("templates/recipe-register.html")
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

func HandleRecipeDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// テンプレートをパースする
		tpl, err := template.ParseFiles("templates/recipe-delete.html")
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
