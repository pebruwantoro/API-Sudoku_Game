package routes

import (
	"app-sudoku/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.POST("/sudoku-games", controllers.SudokuController)
}
