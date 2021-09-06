package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	board "github.com/shotu/craft/board"
)

// type BoardJsonStruct struct {
// 	ID      int     `json:"id"`
// 	Players []int   `json:"players"`
// 	Snakes  [][]int `json:"snakes"`
// 	Ladders [][]int `json:"ladders"`
// }

var (
	boards = map[int]*board.BoardImpl{}
)

func CreateBoard(c echo.Context) error {

	players := []board.Player{}

	// fmt.Println("c params", c.Param("players"))

	player1 := board.NewPlayer(1)
	player2 := board.NewPlayer(2)

	players = append(players, player1)
	players = append(players, player2)
	id := 1

	board := board.NewBoard(board.NewDice(), board.NewSnakes(), board.NewLadders(), players, id)
	fmt.Println("Board is", board)
	if err := c.Bind(board); err != nil {
		return err
	}

	boards[id] = board
	return c.JSON(http.StatusCreated, board)
}

func GetBoard(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	board := boards[id]
	return c.JSON(http.StatusOK, board)
}
