package board

import (
	"fmt"
)

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

func (b *BoardImpl) UpdatePlayerPostion(playerId int, newPos int) ([][]int, [][]int, error) {
	// board := boards[boardID]

	players := b.Players

	snakes := [][]int{}
	ladders := [][]int{}

	// players := boards[boardID].Players
	for _, player := range players {
		fmt.Println("pos ", player.CurrentPosition)
		fmt.Println("id", player.ID)

		if player.ID == playerId {

			fmt.Println("player is ", player)
			newPos := player.CurrentPosition + newPos
			// will not update if new postion > 100

			if newPos <= 100 {
				fmt.Println("update Player Postion.... new pos....", newPos)
				player.CurrentPosition = newPos
			}

			isSnake := true
			isLadder := true
			fmt.Println("isSnake", isSnake)
			fmt.Println("isLadder", isLadder)

			for isSnake || isLadder {

				fmt.Println("Intial is snake is ladder ", isSnake, isLadder)

				snakeEnd, isSnk := b.Snakes.SnakesMap[player.CurrentPosition]
				isSnake = isSnk
				fmt.Println("is snake:", isSnake)

				if isSnake {
					fmt.Println("Bitten by snake")
					snakes = append(snakes, []int{player.CurrentPosition, snakeEnd})
					player.CurrentPosition = snakeEnd
				}

				// time.Sleep(5 * time.Second)

				fmt.Println(" player.  CurrentPosition", player.CurrentPosition)

				ladderEnd, isLad := b.Ladders.LaddersMap[player.CurrentPosition]

				fmt.Println("is isLadder  :", isLadder)
				isLadder = isLad
				if isLadder {

					ladders = append(ladders, []int{player.CurrentPosition, ladderEnd})
					fmt.Println("got the ladder ")
					player.CurrentPosition = ladderEnd
				}
			}
			break
		}
	}

	fmt.Println("Completes update")
	return snakes, ladders, nil
}
