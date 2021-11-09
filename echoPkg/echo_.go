/**
 * Created by GoLand
 * Time: 2021/11/9 11:49 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: echo_.go
 * Desc:
 */
package echoPkg

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"name"`
}

func Main() {
	e := echo.New()

	// middleware 401
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}
	})

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello World")
	})

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		//u.Email = "tzfforyou@163.com"
		//u.Name = "Jerry"

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":8081"))
}
