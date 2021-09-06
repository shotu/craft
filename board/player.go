package board

type Player struct {
	CurrentPosition int
	ID              int
}

func (p *Player) SetCurrPosition(position int) {
	p.CurrentPosition = position
}

func (p *Player) GetPos() int {
	return p.CurrentPosition
}

func NewPlayer(id int) *Player {

	player := &Player{CurrentPosition: 0, ID: id}
	return player
}
