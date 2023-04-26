package Game

type Snake struct {
	body          []XY
	direction     int
	lastDirection int
}

const (
	normal int = iota
	left
	right
	top
	down
)

func NewSnake() *Snake {
	return &Snake{body: []XY{{3, 3}}, direction: normal, lastDirection: normal}
}
