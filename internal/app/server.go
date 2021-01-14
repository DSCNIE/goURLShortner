package app

import (
	"fmt"
	"net/http"
	"os"

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
		e.GET("/home", func(context echo.Context) error {
			return context.String(http.StatusOK, "Hey, I am UP")
		})
		e.GET("/", func(context echo.Context) error {
			return context.Redirect(http.StatusPermanentRedirect, "https://website.dscnie.tech")
		})
		e.POST("/create", router.CreateShortLink)
		e.GET("/:route", router.GetShortenLink)
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
		// e.Logger.Fatal((e.Start(":5000")))
	} else {
		fmt.Println("database did not connect")
	}
}
