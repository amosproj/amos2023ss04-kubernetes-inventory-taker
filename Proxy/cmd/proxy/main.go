package main

import (
	"fmt"
	"net/http"
	"os"

	data "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echoServer := echo.New()

	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	echoServer.GET("/", func(c echo.Context) error {
		return fmt.Errorf("internal server error: %w", c.HTML(http.StatusOK, data.TestData()))
	})

	echoServer.GET("/health", func(c echo.Context) error {
		return fmt.Errorf("internal server error: %w", c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"}))
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	echoServer.Logger.Fatal(echoServer.Start(":" + httpPort))
}
