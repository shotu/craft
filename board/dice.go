package board

import (
	"math/rand"
	"time"
)

type Dice struct{}

// Number
func (d Dice) RollDice() int {
	var min int = 1
	var max int = 6
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(max-min) + min
	return val
}

func NewDice()Dice{
	return Dice{}
}