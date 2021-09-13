package handlersv1

import (
	"encoding/json"
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

// This will beahave as our db
var (
	boards = map[int]*board.BoardImpl{}
)

type snakes struct {
}

type ladders struct{}

type result struct {
	Board   *board.BoardImpl `json:"board"`
	Ladders [][]int          `json:"ladders"`
	Snakes  [][]int          `json:"snakes"`
}

func CreateBoard(c echo.Context) error {

	players := []*board.Player{}

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

func RollDice(c echo.Context) error {

	fmt.Println("rolling the dice")
	boardID, _ := strconv.Atoi(c.Param("board_id"))
	playerID, _ := strconv.Atoi(c.Param("id"))

	//TODO

	board := boards[boardID]
	// fmt.Println(" board player.....................", board.Players.id)

	rolledDiceNumber := board.Dice.RollDice()

	snakes, ladders, err := board.UpdatePlayerPostion(playerID, rolledDiceNumber)

	fmt.Println("snakes..", snakes)
	fmt.Println("ladders..", ladders)
	result := result{
		Board:   board,
		Ladders: ladders,
		Snakes:  snakes,
	}
	if err != nil {
		fmt.Println("error in rolling dice", err)
		return c.JSON(http.StatusBadRequest, err)

	} else {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(result)

	}
	// 	fmt.Println("Updated board is", result)      //
	// 	return c.SetParamNames(http.StatusOK, board) //
	// }
}
