package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gmolveau/webview_echo/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zserge/webview"
	"net/http"
)

func main() {
	e := echo.New()
	go func() {
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))
		api := e.Group("/api")
		{
			api.GET("", controllers.GetExample)
		}

		assetHandler := http.FileServer(rice.MustFindBox("./static").HTTPBox())
		e.GET("/*", echo.WrapHandler(assetHandler))

		e.Logger.Fatal(e.Start(":3333"))
	}()

	webview.Open("My Application Webview",
		"http://localhost:3333", 800, 600, true)
}
