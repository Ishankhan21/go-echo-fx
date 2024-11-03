package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ishankhan21/go-fiber-fx/controllers"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func StartEchoServer(lx fx.Lifecycle, uc *controllers.UserController) *echo.Echo {
	echoServer := echo.New()
	echoServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	echoServer.GET("/profile", uc.GetProfileData)
	echoServer.GET("/listOrders", uc.ListOrders)

	lx.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server ")
			go echoServer.Start(":80")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("On stop hook called")
			return echoServer.Shutdown(ctx)
		},
	})

	return echoServer
}
