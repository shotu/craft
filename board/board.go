package board

import "fmt"

type Board interface {
	NewBoard()
}

type BoardImpl struct {
	Dice    Dice     `json:"dice"`
	Snakes  Snakes   `json:"snakes"`
	Ladders Ladders  `json:"ladders"`
	Players []Player `json:"players"`
	ID      int      `json:"id"`
}

// type BoardJsonStruct struct {
// 	ID      int     `json:"id"`
// 	Players []int   `json:"players"`
// 	Snakes  [][]int `json:"snakes"`
// 	Ladders [][]int `json:"ladders"`
// }

func NewBoard(dice Dice, snakes Snakes, ladders Ladders, players []Player, id int) *BoardImpl {

	fmt.Println("Here")

	boardIns := &BoardImpl{
		Dice:    dice,
		Snakes:  snakes,
		Ladders: ladders,
		Players: players,
		ID:      id,
	}
	fmt.Println("Here boardIns: ", boardIns)

	return boardIns
}
