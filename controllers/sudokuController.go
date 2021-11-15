package controllers

import (
	"app-sudoku/logic"
	"app-sudoku/models"
	"net/http"

	"github.com/labstack/echo"
)

func SudokuController(c echo.Context) error {
	var sudoku models.Sudoku
	c.Bind(&sudoku)
	if logic.SolveSudoku(&sudoku.Input) {
		return c.JSON(http.StatusOK, sudoku.Input)
	} else {
		return c.JSON(http.StatusBadRequest, "The Sudoku can not be solved")
	}
}
