package board

type Player struct {
	currentPosition int
	ID              int
}

func (p *Player) SetCurrPosition(position int) {
	p.currentPosition = position
}

func (p *Player) GetPos() int {
	return p.currentPosition
}

func NewPlayer(id int) Player {

	player := Player{currentPosition: 0, ID: id}
	return player
}
