package board

import "errors"

type Snakes struct {
	SnakesMap map[int]int `json:"snakesMap"`
}

func NewSnakes() Snakes {

	snakeRandMap := getSnakeRandomMap(2, 99)
	snakes := Snakes{
		SnakesMap: snakeRandMap,
	}
	return snakes
}

//TODO combination of ladders and snake should be solvable
// start and end should be different,
func getSnakeRandomMap(start int, end int) map[int]int {

	snakesRandMap := map[int]int{91: 3, 87: 16, 78: 45, 66: 48, 68: 36, 56: 26, 73: 52, 48: 9}
	return snakesRandMap
}

func (s Snakes) GetSnakesEndPos(start int) (int, error) {
	end, ok := s.SnakesMap[start]
	if !ok {
		return 0, errors.New("not a valid start point") // return 0 and error if invalid
	}
	return end, nil
}
