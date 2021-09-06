package board

import "errors"

type Ladders struct {
	LaddersMap map[int]int `json:"laddersMap"`
}

func NewLadders() Ladders {

	laddersRandMap := getRandomMap(2, 99)
	ladders := Ladders{
		LaddersMap: laddersRandMap,
	}
	return ladders
}

//TODO combination of ladders and snake shpuld be solvable
// start and end should be different,
func getRandomMap(start int, end int) map[int]int {

	laddersRandMap := map[int]int{2: 30, 4: 16, 9: 19, 26: 48, 38: 86, 50: 61, 73: 96, 48: 99}
	return laddersRandMap
}

func (l Ladders) GetLadderEndPos(start int) (int, error) {
	end, ok := l.LaddersMap[start]
	if !ok {
		return 0, errors.New("Not a valid start point") // return 0 and error if invalid
	}
	return end, nil
}
