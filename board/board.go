package board

import "fmt"

type Board interface {
	NewBoard()
}

type BoardImpl struct {
	Dice    Dice      `json:"dice"`
	Snakes  Snakes    `json:"snakes"`
	Ladders Ladders   `json:"ladders"`
	Players []*Player `json:"players"`
	ID      int       `json:"id"`
	Start   int       `json:"start"`
	End     int       `json:"end"`
}

// type BoardJsonStruct struct {
// 	ID      int     `json:"id"`
// 	Players []int   `json:"players"`
// 	Snakes  [][]int `json:"snakes"`
// 	Ladders [][]int `json:"ladders"`
// }

func NewBoard(dice Dice, snakes Snakes, ladders Ladders, players []*Player, id int) *BoardImpl {

	fmt.Println("Here")

	boardIns := &BoardImpl{
		Dice:    dice,
		Snakes:  snakes,
		Ladders: ladders,
		Players: players,
		ID:      id,
		Start:   1,
		End:     100,
	}
	fmt.Println("Here boardIns: ", boardIns)

	return boardIns
}

func (b *BoardImpl) UpdatePlayerPostion(playerId int, newPos int) error {
	// board := boards[boardID]

	players := b.Players
	// players := boards[boardID].Players
	for _, player := range players {
		fmt.Println("pos ", player.CurrentPosition)
		fmt.Println("id", player.ID)

		if player.ID == playerId {

			fmt.Println("player is ", player)
			newPos := player.CurrentPosition + newPos
			// will not update if new postion > 100

			fmt.Println("updatePlayerPostion.... new pos....", newPos)
			if newPos <= 100 {
				player.CurrentPosition = newPos
			}

			var isSnake bool
			var isLadder bool

			for isSnake && isLadder {

				fmt.Println("checking if snake or ladder .... new pos....", newPos)
				snakeStart, isSnake := b.Snakes.SnakesMap[player.CurrentPosition]

				if isSnake {
					fmt.Println("Bitten by snake")
					player.CurrentPosition = snakeStart
				}

				ladderEnd, isLadder := b.Ladders.LaddersMap[player.CurrentPosition]

				if isLadder {
					fmt.Println("got the ladder ")
					player.CurrentPosition = ladderEnd
				}
			}

			break
		}

	}

	fmt.Println("Completes update")

	return nil
}
