package main

import (
	"app-sudoku/config"
	"app-sudoku/middlewares"
	"app-sudoku/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	config.InitPort()
	middlewares.LogMiddlewares((app))
	routes.New(app)
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
