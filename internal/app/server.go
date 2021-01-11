package app

import (
	"fmt"

	"github.com/iresharma/REST1/internal/pkg/database"
	"github.com/iresharma/REST1/internal/pkg/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run - runs the server
func Run() {
	e := echo.New()
	err := database.Conn()
	if err {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))
		e.POST("/create", router.CreateShortLink)
		e.GET("/short/:route", router.GetShortenLink)
		e.Logger.Fatal(e.Start(":5000"))
	} else {
		fmt.Println("database did not connect")
	}
}
